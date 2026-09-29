#!/usr/bin/env python3
"""Scrape the official Brightspace HTTP routing table into a CSV file.

Install dependencies with ``python -m pip install -r docs/routes/requirements.txt``.
Run ``python docs/routes/scrape_routes.py`` from any directory. When a route
has multiple JSON parameter types, they are separated by semicolons.
"""

from __future__ import annotations

import argparse
from concurrent.futures import ThreadPoolExecutor
import csv
import os
from pathlib import Path
import re
import sys
import tempfile
from urllib.error import URLError
from urllib.parse import unquote, urljoin, urlsplit, urlunsplit
from urllib.request import Request, urlopen

from bs4 import BeautifulSoup


ROUTING_TABLE = "https://docs.valence.desire2learn.com/http-routingtable.html"
OUTPUT = Path(__file__).resolve().parent / "routes.csv"
ROUTE_LABEL = re.compile(r"^(?P<route>/.*?)\s+\[(?P<method>[A-Z]+)\]$")
COLUMNS = (
    "category",
    "method",
    "route",
    "json_parameters",
    "return_type",
    "active_versions",
    "deprecated_versions",
    "obsolete_versions",
    "documentation_url",
)


def fetch(url: str) -> BeautifulSoup:
    request = Request(url, headers={"User-Agent": "GoValence-route-scraper/1.0"})
    with urlopen(request, timeout=30) as response:
        return BeautifulSoup(response.read(), "html.parser")


def clean_text(element) -> str:
    return " ".join(element.get_text(" ", strip=True).split())


def index_routes(index_url: str) -> list[dict[str, str]]:
    page = fetch(index_url)
    index_host = urlsplit(index_url).netloc
    category = ""
    routes = []
    for row in page.select("table.indextable tr"):
        heading = row.select_one("strong") if "cap" in row.get("class", []) else None
        if heading:
            category = clean_text(heading)
            continue
        link = row.select_one("a[href]")
        if link is None:
            continue
        label = ROUTE_LABEL.fullmatch(clean_text(link))
        if label is None:
            continue
        direct_url = urljoin(index_url, link["href"])
        parsed = urlsplit(direct_url)
        if parsed.netloc != index_host or not parsed.path.endswith(".html") or not parsed.fragment:
            raise ValueError(f"Unexpected route documentation link: {direct_url}")
        routes.append(
            {
                "category": category,
                "method": label["method"],
                "route": label["route"],
                "documentation_url": direct_url,
            }
        )
    if not routes:
        raise ValueError(f"No routes found at {index_url}")
    return routes


def field_lists(body) -> dict[str, object]:
    result = {}
    for label in body.select("dl.field-list > dt"):
        name = clean_text(label).rstrip(" :").lower()
        value = label.find_next_sibling("dd")
        if value is not None:
            result[name] = value
    return result


def json_parameters(field_list: dict[str, object]) -> str:
    body = field_list.get("json parameters")
    if body is None:
        return ""
    result = []
    entries = body.select("li") or [body]
    for entry in entries:
        strong = entry.find("strong")
        name = clean_text(strong) if strong else ""
        type_link = entry.select_one("a.reference.internal[title]")
        model_type = type_link.get("title", "") if type_link else ""
        if not model_type:
            parenthesized = re.search(r"\(([^)]+)\)", clean_text(entry))
            model_type = parenthesized.group(1).strip() if parenthesized else name
        if model_type and model_type not in result:
            result.append(model_type)
    return "; ".join(result)


def version_info(field_list: dict[str, object]) -> dict[str, list[str]]:
    grouped: dict[str, list[str]] = {key: [] for key in ("active", "deprecated", "obsolete")}
    body = field_list.get("api versions")
    if body is None:
        return grouped
    for entry in body.select("li") or [body]:
        text = clean_text(entry)
        strong = entry.find("strong")
        version_range = clean_text(strong) if strong else text.split("–", 1)[0].strip()
        note = text[len(version_range):].lstrip(" –—:-")
        lower = note.lower()
        if "obsolete" in lower:
            status = "obsolete"
        elif "deprecated" in lower:
            status = "deprecated"
        elif "not supported" in lower or "unsupported" in lower:
            continue
        else:
            status = "active"
        grouped[status].append(version_range)
    return grouped


def return_type(body) -> str:
    for paragraph in body.find_all("p", recursive=False):
        lead = paragraph.find("strong", recursive=False)
        if not lead or clean_text(lead).rstrip(" .:").lower() != "return":
            continue
        for link in paragraph.select("a.reference.internal"):
            fragment = urlsplit(link.get("href", "")).fragment.lower()
            if fragment.startswith(("get--", "post--", "put--", "delete--")):
                continue
            type_name = link.get("title") or clean_text(link)
            if type_name:
                return type_name
        text = clean_text(paragraph)
        qualified = re.search(r"\b[A-Z][A-Za-z0-9_]*\.[A-Z][A-Za-z0-9_]*\b", text)
        if qualified:
            return qualified.group()
        for marker, type_name in (
            (r"\bJSON boolean\b", "boolean"),
            (r"\b(?:application/xml|XML document)\b", "XML"),
            (r"\btext/plain\b", "text/plain"),
            (r"\bCSV\b", "CSV"),
            (r"\bPDF\b", "PDF"),
            (r"\b(?:image|images)\b", "image"),
            (r"\b(?:file|binary|stream)\b", "file"),
            (r"\bJSON array\b|\b(?:array|list) of\b", "JSON array"),
            (r"\bJSON (?:block|object)\b", "JSON object"),
            (r"\b(?:does not return a response body|empty response body|no response body)\b", "none"),
        ):
            if re.search(marker, text, re.I):
                return type_name
    return ""


def route_definitions(page_url: str) -> dict[str, object]:
    page = fetch(page_url)
    definitions = {}
    for definition in page.select("dl.http"):
        heading = definition.find("dt", recursive=False)
        body = definition.find("dd", recursive=False)
        if heading and body and heading.get("id"):
            definitions[heading["id"]] = body
    return definitions


def scrape(index_url: str) -> list[dict[str, str]]:
    routes = index_routes(index_url)
    pages = sorted({urlunsplit(urlsplit(route["documentation_url"])._replace(fragment="")) for route in routes})
    with ThreadPoolExecutor(max_workers=4) as pool:
        definitions = dict(zip(pages, pool.map(route_definitions, pages)))
    rows = []
    missing = []
    for route in routes:
        parsed = urlsplit(route["documentation_url"])
        page_url = urlunsplit(parsed._replace(fragment=""))
        body = definitions[page_url].get(unquote(parsed.fragment))
        if body is None:
            missing.append(route["documentation_url"])
            continue
        fields = field_lists(body)
        versions = version_info(fields)
        rows.append(
            {
                **route,
                "json_parameters": json_parameters(fields),
                "return_type": return_type(body),
                "active_versions": "; ".join(versions["active"]),
                "deprecated_versions": "; ".join(versions["deprecated"]),
                "obsolete_versions": "; ".join(versions["obsolete"]),
            }
        )
    if missing:
        sample = ", ".join(missing[:3])
        raise ValueError(f"{len(missing)} routing-table links lack route definitions: {sample}")
    return rows


def write_csv(rows: list[dict[str, str]], output: Path) -> None:
    output.parent.mkdir(parents=True, exist_ok=True)
    temp_name = None
    try:
        with tempfile.NamedTemporaryFile("w", encoding="utf-8", newline="", dir=output.parent, delete=False) as stream:
            temp_name = stream.name
            writer = csv.DictWriter(stream, fieldnames=COLUMNS)
            writer.writeheader()
            writer.writerows(rows)
        os.replace(temp_name, output)
    finally:
        if temp_name and os.path.exists(temp_name):
            os.unlink(temp_name)


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--index-url", default=ROUTING_TABLE)
    parser.add_argument("--output", type=Path, default=OUTPUT)
    args = parser.parse_args()
    try:
        rows = scrape(args.index_url)
        write_csv(rows, args.output)
    except (OSError, URLError, ValueError) as exc:
        print(f"Scrape failed: {exc}", file=sys.stderr)
        return 1
    print(f"Wrote {len(rows)} routes to {args.output}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
