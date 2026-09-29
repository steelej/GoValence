#!/usr/bin/env python3
"""Compare routes used by Go API methods with docs/routes/routes.csv.

Run ``python docs/routes/compose_routes.py`` from any directory. Only routes
implemented by this client are checked; the documentation lists many more.
Exit status is 1 when an inconsistency or an unreadable route is found.
"""

from __future__ import annotations

import argparse
import csv
from dataclasses import dataclass
import difflib
from pathlib import Path
import re
import sys


ROOT = Path(__file__).resolve().parents[2]
CSV_PATH = Path(__file__).with_name("routes.csv")
METHOD = re.compile(r"^func \(c \*Client\) (?P<name>\w+)\(", re.M)
REQUEST = re.compile(r"\bc\.(getRaw|get|postJSON|postMultipartFile|putJSON|delete)\s*\(")
PATH_BUILDER = re.compile(r"^c\.(lpPath|lpUnstablePath|lePath|leUnstablePath|basPath)\s*\(")
GO_STRING = re.compile(r'^"(?:[^"\\]|\\.)*"$')
GO_PLACEHOLDER = re.compile(r"%[sd]")
DOC_PLACEHOLDER = re.compile(r"\([^)]+\)")
HTTP_METHOD = {
    "get": "GET", "getRaw": "GET", "postJSON": "POST",
    "postMultipartFile": "POST", "putJSON": "PUT", "delete": "DELETE",
}
PREFIX = {
    "lpPath": "/d2l/api/lp/(version)/",
    "lpUnstablePath": "/d2l/api/lp/unstable/",
    "lePath": "/d2l/api/le/(version)/",
    "leUnstablePath": "/d2l/api/le/unstable/",
    "basPath": "/d2l/api/bas/(version)/",
}


@dataclass(frozen=True)
class GoRoute:
    method: str
    route: str
    function: str
    file: Path
    line: int


def first_argument(source: str, open_paren: int) -> str:
    """Read a Go call's first argument, allowing nested calls and strings."""
    depth = 1
    in_string = False
    escaped = False
    start = open_paren + 1
    for pos in range(start, len(source)):
        char = source[pos]
        if in_string:
            if escaped:
                escaped = False
            elif char == "\\":
                escaped = True
            elif char == '"':
                in_string = False
        elif char == '"':
            in_string = True
        elif char == "(":
            depth += 1
        elif char == ")":
            depth -= 1
            if depth == 0:
                return source[start:pos].strip()
        elif char == "," and depth == 1:
            return source[start:pos].strip()
    raise ValueError("unterminated call")


def go_string(value: str) -> str:
    if not GO_STRING.fullmatch(value):
        raise ValueError(f"expected a string literal, got {value[:80]!r}")
    # Path literals in this repository use ordinary Go double-quoted strings.
    import json
    return json.loads(value)


def path_from_expression(expression: str) -> str:
    builder = PATH_BUILDER.match(expression)
    if builder:
        value = first_argument(expression, builder.end() - 1)
        return PREFIX[builder.group(1)] + GO_PLACEHOLDER.sub("(value)", go_string(value))
    if GO_STRING.fullmatch(expression):
        return go_string(expression)
    raise ValueError(f"unsupported path expression {expression[:100]!r}")


def go_routes(root: Path) -> tuple[list[GoRoute], list[str]]:
    routes: list[GoRoute] = []
    errors: list[str] = []
    for file in sorted(root.glob("api_*.go")):
        if file.name.endswith("_test.go"):
            continue
        source = file.read_text(encoding="utf-8")
        functions = list(METHOD.finditer(source))
        for index, function in enumerate(functions):
            end = functions[index + 1].start() if index + 1 < len(functions) else len(source)
            body = source[function.start():end]
            for call in REQUEST.finditer(body):
                offset = function.start() + call.start()
                line = source.count("\n", 0, offset) + 1
                try:
                    expression = first_argument(body, call.end() - 1)
                    route = path_from_expression(expression)
                except ValueError as exc:
                    errors.append(f"{file.name}:{line} {function.group('name')}: {exc}")
                    continue
                routes.append(GoRoute(HTTP_METHOD[call.group(1)], route,
                                      function.group("name"), file, line))
    return routes, errors


def canonical(route: str) -> str:
    # Go's format verbs and documentation's parameter names describe the same
    # varying path segments. Keep spelling, case and trailing slash intact.
    route = DOC_PLACEHOLDER.sub("(value)", route)
    return GO_PLACEHOLDER.sub("(value)", route)


def equivalent(go: str, documented: str) -> bool:
    if canonical(go) == canonical(documented):
        return True
    # The unstable helpers fix the version segment to the literal "unstable".
    return canonical(re.sub(r"(/api/(?:lp|le)/)unstable(?=/)", r"\1(version)", go)) == canonical(documented)


def load_documentation(csv_path: Path) -> list[dict[str, str]]:
    with csv_path.open(newline="", encoding="utf-8-sig") as handle:
        reader = csv.DictReader(handle)
        if not {"method", "route", "documentation_url"}.issubset(reader.fieldnames or []):
            raise ValueError("CSV needs method, route and documentation_url columns")
        rows = list(reader)
    if not rows:
        raise ValueError("CSV has no routes")
    return rows


def compare(routes: list[GoRoute], docs: list[dict[str, str]]) -> list[str]:
    issues = []
    for route in routes:
        matches = [row for row in docs if equivalent(route.route, row["route"])]
        if any(row["method"] == route.method for row in matches):
            continue
        location = f"{route.file.name}:{route.line} {route.function}"
        if matches:
            details = "; ".join(f"{row['method']} {row['route']} {row['documentation_url']}" for row in matches)
            issues.append(f"{location}: {route.method} {route.route}\n  method differs; documented: {details}")
            continue
        candidates = [row for row in docs if row["method"] == route.method]
        closest = difflib.get_close_matches(canonical(route.route),
                                            [canonical(row["route"]) for row in candidates],
                                            n=3, cutoff=0.65)
        suggestions = []
        for candidate in closest:
            row = next(row for row in candidates if canonical(row["route"]) == candidate)
            suggestions.append(f"{row['method']} {row['route']} {row['documentation_url']}")
        detail = "\n  closest: " + "; ".join(suggestions) if suggestions else ""
        issues.append(f"{location}: {route.method} {route.route}\n  route missing from CSV{detail}")
    return issues


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--source", type=Path, default=ROOT, help="directory with api_*.go files")
    parser.add_argument("--csv", type=Path, default=CSV_PATH, help="scraped route CSV")
    args = parser.parse_args()
    try:
        docs = load_documentation(args.csv)
        routes, errors = go_routes(args.source)
        if not routes and not errors:
            raise ValueError(f"No Go HTTP calls found in {args.source}")
    except (OSError, ValueError) as exc:
        print(f"compose_routes: {exc}", file=sys.stderr)
        return 2
    issues = compare(routes, docs)
    for issue in errors + issues:
        print(issue)
    print(f"Checked {len(routes)} Go HTTP calls against {len(docs)} documented routes: "
          f"{len(routes) - len(issues)} matched, {len(issues)} mismatched, "
          f"{len(errors)} unreadable.")
    return 1 if issues or errors else 0


if __name__ == "__main__":
    sys.exit(main())
