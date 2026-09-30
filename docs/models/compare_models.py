#!/usr/bin/env python3
"""Compare the JSON fields in models.go with the saved D2L model definitions.

Run ``python docs/models/compare_models.py``. Intentional differences are recorded
in overrides.json and remain visible in the output. Exit status is 1 when there
are unacknowledged differences, stale overrides, or Go structs without an
identified documentation counterpart.
The documentation files are produced by scrape_models.py; this script is offline.
"""

from __future__ import annotations

import argparse
from dataclasses import dataclass, field
import json
from pathlib import Path
import re


HERE = Path(__file__).resolve().parent
GO_FILE = HERE.parents[1] / "models.go"

# Go names that differ from the published model names. Nested locations use
# "Group/Model#Field#Subfield". Keep this table explicit: a guessed match can
# silently compare the wrong API contract.
ALIASES: dict[str, str] = {
    "PagingInfo": "Api/PagedResultSet#PagingInfo",
    "OrganizationInfo": "Org/Organization",
    "OrgUnitTypePermissions": "OrgUnit/Permissions",
    "AccessInfo": "Enrollment/MyOrgUnitInfo#Access",
    "OrgUnitUserInfo": "User/User",
    "UserEnrollmentData": "Enrollment/UserOrgUnit",
    "GroupData": "Group/GroupData@1",
    "Group": "Group/GroupData@2",
    "GroupCategoryData": "Group/GroupCategoryData@1",
    "GroupCategory": "Group/GroupCategoryData@2",
    "SectionData": "Section/SectionData@1",
    "Section": "Section/SectionData@2",
    "SectionPropertyData": "Section/SectionSettingsData@2",
    "GradeSchemeRange": "Grade/GradeScheme@1#Ranges",
    "GradeScheme": "Grade/GradeScheme@1",
    "GradeCategory": "Grade/GradeObjectCategory",
    "GradeValueEntry": "Grade/UserGradeValue",
    "FinalGradeValueEntry": "Grade/UserGradeValue",
    "NewsFile": "News/NewsItem#Attachments",
    "QuizAttemptsAllowed": "Quiz/QuizReadData#AttemptsAllowed",
    "QuizLateSubmissionInfo": "Quiz/QuizReadData#LateSubmissionInfo",
    "QuizIPRange": "Quiz/QuizReadData#RestrictIPAddressRange",
    "TimeLimit": "Quiz/QuizReadData#SubmissionTimeLimit",
    "Instructions": "Quiz/QuizReadData#Instructions",
    "Description": "Quiz/QuizReadData#Description",
    "QuizQuestion": "Quiz/QuestionData",
    "QuizSpecialAccessData": "Quiz/SpecialAccessData",
    "QuizSpecialAccessTimeLimit": "Quiz/SpecialAccessData#SubmissionTimeLimit",
    "ContentModule": "ToC/TableOfContents#Modules#Modules",
    "ContentTopic": "ToC/TableOfContents#Modules#Topics",
    "RubricLevel": "Rubric/Level",
    "RubricCell": "Rubric/CriteriaGroup#Criteria#Cells",
    "RubricCriterion": "Rubric/CriteriaGroup#Criteria",
    "RubricCriteriaGroup": "Rubric/CriteriaGroup",
    "RubricOverallLevel": "Rubric/OverallLevel",
    "DropboxRubric": "Rubric/Rubric",
    "DropboxAssessment": "Dropbox/DropboxFolder#Assessment",
    "DropboxAvailability": "Dropbox/DropboxFolder#Availability",
    "SubmissionFile": "Dropbox/DropboxFolder#Attachments",
    "DropboxEntity": "Dropbox/Entity",
    "DropboxFeedback": "Dropbox/DropboxFeedbackOut",
    "DropboxLink": "Dropbox/DropboxFeedbackOut#Links",
    "DropboxSubmitter": "Dropbox/EntityDropbox#Submissions#SubmittedBy",
    "DropboxSubmissionFile": "Dropbox/EntityDropbox#Submissions#Files",
    "DropboxSubmissionEntry": "Dropbox/EntityDropbox#Submissions",
    "UserSubmissions": "Dropbox/EntityDropbox",
    "Survey": "Surveys/SurveyReadData",
    "SurveyUserResponses": "Surveys/SurveyReadData#UserResponses",
    "SurveyAttempt": "Surveys/SurveyAttemptData",
    "LTILink": "LTI/LtiLinkData",
    "LTIAdvantageLink": "LTI/LTIAdvantageLinkData",
    "LTIAdvantageQuicklink": "LTI/LtiQuickLinkData",
    "LTICustomParameter": "LTI/CustomParameter",
    "LTIToolProvider": "LTI/LtiToolProviderData",
    "LTISharingData": "LTI/SharingRuleWithDescendantTypes",
    "LTIDeploymentSharingData": "LTI/OrgUnitSharingRuleData",
    "LTIAdvantageCreateSharingRuleData": "LTI/CreateSharingRuleData",
    "ToolInfo": "Tools/OrgUnitInformation",
    "IntelligentAgent": "IntelligentAgents/AgentData",
    "IntelligentAgentSchedule": "IntelligentAgents/ScheduleData",
    "IntelligentAgentAction": "IntelligentAgents/ActionData",
    "IntelligentAgentEmailAction": "IntelligentAgents/EmailAction",
    "IntelligentAgentEnrollAction": "IntelligentAgents/EnrollmentAction",
    "IntelligentAgentCondition": "IntelligentAgents/ConditionData",
    "IntelligentAgentDateCondition": "IntelligentAgents/DateCondition",
    "IntelligentAgentReleaseCondition": "IntelligentAgents/ReleaseCondition",
    "ConfigVariableValue": "ConfigVariable/OrgUnitValue",
    "ReleaseConditionsData": "ReleaseConditions/ConditionsData",
    "CourseImportJobData": "Course/CreateCopyJobResponse",
    "CourseImportJobStatus": "Course/GetImportJobResponse",
    "IssuedBadge": "Awards/IssuedAward",
    "IssuedAwardShare": "Awards/IssuedAward#Share",
    "AwardExpiryCalculation": "Awards/ExpiryCalculation",
    "AwardExpiryNotification": "Awards/ExpiryNotification",
    "AwardFileData": "Awards/FileData",
    "OrgToolInfo":"Tools/OrgInformation",
    "CreateSectionSettingsData":"Section/SectionSettingsData@1",
    "UpdateSectionSettingsData":"Section/SectionSettingsData@3",
}


@dataclass
class Shape:
    kinds: set[str] = field(default_factory=set)
    fields: dict[str, "Shape"] = field(default_factory=dict)

    def merge(self, other: "Shape") -> None:
        self.kinds.update(other.kinds)
        for name, child in other.fields.items():
            if name in self.fields:
                self.fields[name].merge(child)
            else:
                self.fields[name] = child


TOKEN = re.compile(
    r'//[^\n]*|"(?:\\.|[^"\\])*"|<[^>\n]*>|\.\.\.|[{}\[\]:,|]|[A-Za-z_][\w.:-]*|-?\d+(?:\.\d+)?|\S'
)


def tokenize(text: str) -> list[str]:
    # One published model accidentally puts an entire field inside quotes.
    text = re.sub(
        r'"([A-Za-z_]\w*):\s*(null\|\[[^"\n]+\])"',
        lambda match: f'"{match[1]}": {match[2]}',
        text,
    )
    return [m.group() for m in TOKEN.finditer(text) if not m.group().startswith("//")]


class Parser:
    def __init__(self, tokens: list[str]):
        self.tokens = tokens
        self.pos = 0

    def peek(self) -> str:
        return self.tokens[self.pos] if self.pos < len(self.tokens) else ""

    def pop(self) -> str:
        value = self.peek()
        self.pos += bool(value)
        return value

    def value(self) -> Shape:
        result = self.primary()
        while self.peek() == "|":
            self.pop()
            result.merge(self.primary())
        return result

    def primary(self) -> Shape:
        token = self.pop()
        if token == "{":
            result = Shape({"object"})
            while self.peek() and self.peek() != "}":
                key = self.pop()
                if key.startswith('"') and self.peek() == ":":
                    self.pop()
                    name = key[1:-1]
                    result.fields[name] = self.value()
                elif key.startswith("<composite:"):
                    pass
                if self.peek() == ",":
                    self.pop()
            if self.peek() == "}":
                self.pop()
            return result
        if token == "[":
            item = Shape()
            while self.peek() and self.peek() != "]":
                if self.peek() in (",", "..."):
                    self.pop()
                    continue
                start = self.pos
                item.merge(self.value())
                if self.pos == start:
                    self.pop()
            if self.peek() == "]":
                self.pop()
            return Shape({"array:" + "/".join(sorted(item.kinds or {"unknown"}))}, item.fields)
        if token.startswith("<"):
            base = token[1:-1].split(":", 1)[0].lower()
            return Shape({
                {"boolean": "bool", "bool": "bool", "composite": "object", "int": "number",
                 "integer": "number", "d2lid": "number", "guid": "string",
                 "utcdatetime": "string", "number": "number",
                 "string": "string", "orgunittypeid": "number"}.get(
                    base, "object" if "." in base else "unknown"
                )
            })
        if token == "null":
            return Shape({"null"})
        if token in ("true", "false"):
            return Shape({"bool"})
        if token.startswith('"'):
            return Shape({"object" if token.startswith('"{composite:') else "string"})
        if re.fullmatch(r"-?\d+(?:\.\d+)?", token):
            return Shape({"number"})
        return Shape({"unknown"})


def read_document(path: Path) -> tuple[str, list[Shape]]:
    content = path.read_text(encoding="utf-8")
    url, separator, body = content.partition("\n\n")
    if not separator:
        raise ValueError(f"Invalid model file: {path}")
    parser = Parser(tokenize(body))
    variants = []
    while parser.peek():
        if parser.peek() == "{":
            obj = parser.value()
            if obj.fields:
                variants.append(obj)
        else:
            parser.pop()
    return url, variants


@dataclass
class GoField:
    name: str
    type: str
    json_string: bool
    line: int


@dataclass
class GoStruct:
    name: str
    fields: dict[str, GoField] = field(default_factory=dict)
    embedded: list[str] = field(default_factory=list)


STRUCT_START = re.compile(r"^type (\w+)(?:\[[^]]+\])? struct \{")
FIELD_LINE = re.compile(r'^\s*(\w+)\s+([*\[\]\w.]+)\s+`json:"([^\"]+)"`')
EMBEDDED_LINE = re.compile(r"^\s+([A-Za-z_]\w*)\s*$")


def read_go(path: Path) -> dict[str, GoStruct]:
    result = {}
    current = None
    for line_no, line in enumerate(path.read_text(encoding="utf-8").splitlines(), 1):
        match = STRUCT_START.match(line)
        if match:
            current = GoStruct(match[1])
            result[current.name] = current
            continue
        if current is None:
            continue
        if line == "}":
            current = None
            continue
        match = FIELD_LINE.match(line)
        if match:
            go_name, go_type, tag = match.groups()
            json_name, *options = tag.split(",")
            if json_name != "-":
                current.fields[json_name or go_name] = GoField(
                    json_name or go_name, go_type, "string" in options, line_no
                )
            continue
        match = EMBEDDED_LINE.match(line)
        if match:
            current.embedded.append(match[1])
        elif line.strip() and not line.lstrip().startswith("//"):
            raise ValueError(f"Cannot parse Go struct field at {path}:{line_no}: {line}")
    for struct in result.values():
        for embedded in struct.embedded:
            if embedded not in result:
                raise ValueError(f"Unknown embedded struct {embedded} in {struct.name}")
            for name, go_field in result[embedded].fields.items():
                struct.fields.setdefault(name, go_field)
    return result


def go_kinds(go_type: str, as_string: bool, structs: dict[str, GoStruct]) -> set[str]:
    if as_string:
        return {"string"}
    typ = go_type.lstrip("*")
    if typ == "[]json.RawMessage":
        return {"array:object", "array:number", "array:string", "array:bool"}
    if typ == "[]T":
        return {"array:object", "array:number", "array:string", "array:bool"}
    if typ.startswith("[]"):
        element = go_kinds(typ[2:], False, structs)
        return {"array:" + "/".join(sorted(element))}
    if typ in ("int", "int64", "int32", "float64", "float32", "uint", "uint64"):
        return {"number"}
    if typ == "bool":
        return {"bool"}
    if typ == "string":
        return {"string"}
    if typ == "NumberOrString":
        return {"number", "string"}
    if typ in ("GradeSchemeEntry", "GradeUserRef", "GradeValueData"):
        return {"object"}
    if typ == "json.RawMessage":
        return {"unknown", "object", "array:object"}
    if typ in structs:
        return {"object"}
    if typ == "T":
        return {"unknown"}
    if typ.startswith("map["):
        return {"object"}
    # Named scalar enums in models.go use an integer underlying type.
    return {"number"}


def find_document(name: str, docs: dict[str, tuple[str, list[Shape]]]) -> str | None:
    if name in ALIASES:
        return ALIASES[name]
    matches = [key for key in docs if key.rsplit("/", 1)[-1] == name]
    return matches[0] if len(matches) == 1 else None


@dataclass
class Difference:
    kind: str
    field: str
    message: str


@dataclass
class Override:
    model: str
    document: str
    field: str
    kind: str
    documented_kinds: set[str]
    go_type: str | None
    json_string: bool
    reason: str

    @property
    def key(self) -> tuple[str, str, str]:
        return self.model, self.field, self.kind

    def matches(self, document: str, struct: GoStruct, shape: Shape) -> bool:
        doc_field = shape.fields.get(self.field)
        go_field = struct.fields.get(self.field)
        return (
            self.document == document
            and self.documented_kinds == (doc_field.kinds if doc_field else set())
            and self.go_type == (go_field.type if go_field else None)
            and self.json_string == (go_field.json_string if go_field else False)
        )


def read_overrides(path: Path) -> dict[tuple[str, str, str], Override]:
    entries = json.loads(path.read_text(encoding="utf-8"))
    if not isinstance(entries, list):
        raise ValueError(f"{path}: overrides must be a JSON array")
    result = {}
    required = {"model", "document", "field", "kind", "documented_kinds", "go_type", "reason"}
    for index, entry in enumerate(entries, 1):
        label = f"{path}: override {index}"
        if not isinstance(entry, dict) or not required <= entry.keys():
            raise ValueError(f"{label}: requires {', '.join(sorted(required))}")
        if entry.keys() - required - {"json_string"}:
            raise ValueError(f"{label}: unknown override properties")
        if any(not isinstance(entry[key], str) or not entry[key].strip()
               for key in ("model", "document", "field", "kind", "reason")):
            raise ValueError(f"{label}: names and reason must be nonempty strings")
        if entry["kind"] not in {"type", "nullable", "missing", "extra"}:
            raise ValueError(f"{label}: kind must be type, nullable, missing, or extra")
        kinds = entry["documented_kinds"]
        if not isinstance(kinds, list) or any(not isinstance(kind, str) or not kind for kind in kinds):
            raise ValueError(f"{label}: documented_kinds must be an array of strings")
        if entry["go_type"] is not None and (
            not isinstance(entry["go_type"], str) or not entry["go_type"].strip()
        ):
            raise ValueError(f"{label}: go_type must be a nonempty string or null")
        if not isinstance(entry.get("json_string", False), bool):
            raise ValueError(f"{label}: json_string must be a boolean")
        override = Override(**{**entry, "documented_kinds": set(kinds),
                               "json_string": entry.get("json_string", False)})
        if override.key in result:
            raise ValueError(f"{label}: duplicate override for {override.key}")
        result[override.key] = override
    return result


def compare(
    struct: GoStruct, shape: Shape, structs: dict[str, GoStruct]
) -> list[Difference]:
    differences = []
    for name, doc_field in shape.fields.items():
        go_field = struct.fields.get(name)
        if go_field is None:
            differences.append(Difference(
                "missing", name, f"  missing Go field: {name} ({'/'.join(sorted(doc_field.kinds))})"
            ))
            continue
        expected = doc_field.kinds - {"null", "unknown", "array:unknown"}
        actual = go_kinds(go_field.type, go_field.json_string, structs)
        if expected and not expected.intersection(actual):
            differences.append(Difference(
                "type", name,
                f"  type {name} (models.go:{go_field.line}): "
                f"Go {go_field.type} emits {'/'.join(sorted(actual))}; "
                f"docs {'/'.join(sorted(expected))}"
            ))
        if (
            "null" in doc_field.kinds
            and not go_field.type.startswith(("*", "[]", "map["))
            and go_field.type != "json.RawMessage"
        ):
            differences.append(Difference(
                "nullable", name,
                f"  nullable {name} (models.go:{go_field.line}): "
                f"docs allow null; Go {go_field.type} cannot represent it"
            ))
    for name, go_field in struct.fields.items():
        if name not in shape.fields:
            differences.append(Difference(
                "extra", name, f"  extra Go field: {name} (models.go:{go_field.line})"
            ))
    return differences


def main(argv: list[str] | None = None) -> int:
    arg_parser = argparse.ArgumentParser(description=__doc__)
    arg_parser.add_argument("--go-file", type=Path, default=GO_FILE)
    arg_parser.add_argument("--models-dir", type=Path, default=HERE)
    arg_parser.add_argument(
        "--overrides", type=Path,
        help="Override registry (default: overrides.json in --models-dir, if present)",
    )
    args = arg_parser.parse_args(argv)
    override_path = args.overrides or args.models_dir / "overrides.json"
    try:
        overrides = read_overrides(override_path) if args.overrides or override_path.exists() else {}
    except (OSError, ValueError) as error:
        arg_parser.error(str(error))
    structs = read_go(args.go_file)
    docs = {
        str(path.relative_to(args.models_dir)): read_document(path)
        for path in sorted(args.models_dir.glob("*/*"))
        if path.is_file() and path.parent.name != "__pycache__"
    }
    different = 0
    unmatched = 0
    matched = 0
    intentional = 0
    used_overrides = set()
    for name, struct in structs.items():
        match = find_document(name, docs)
        if match is None:
            print(f"{name}: no unambiguous standalone documentation model")
            unmatched += 1
            continue
        model_match, _, field_path = match.partition("#")
        key, _, variant_text = model_match.partition("@")
        if key not in docs:
            raise ValueError(f"Missing documentation model {key} for {name}")
        url, variants = docs[key]
        if not variants:
            print(f"{name}: no parseable object in {key}")
            unmatched += 1
            continue
        if variant_text:
            variant_number = int(variant_text)
            if not 1 <= variant_number <= len(variants):
                raise ValueError(f"Invalid variant {match} for {name}")
            shape = variants[variant_number - 1]
        else:
            shape = Shape({"object"})
            for variant in variants:
                shape.merge(variant)
        for part in filter(None, field_path.split("#")):
            if part not in shape.fields:
                raise ValueError(f"Missing documentation field {part} in {match} for {name}")
            shape = shape.fields[part]
        issues = compare(struct, shape, structs)
        matched += 1
        if issues:
            print(f"{name} <> {key} ({url})")
            unacknowledged = False
            for issue in issues:
                override = overrides.get((name, issue.field, issue.kind))
                if override and override.matches(match, struct, shape):
                    used_overrides.add(override.key)
                    intentional += 1
                    print(f"{issue.message} [intentional override]")
                    print(f"    reason: {override.reason}")
                else:
                    unacknowledged = True
                    print(issue.message)
            different += int(unacknowledged)
    stale = overrides.keys() - used_overrides
    for model, field_name, kind in sorted(stale):
        print(f"Stale override: {model}.{field_name} ({kind}) no longer matches an observed difference; "
              "review or remove it")
    print(
        f"Checked {len(structs)} structs: {matched} matched, "
        f"{different} with unacknowledged differences, {unmatched} without a counterpart, "
        f"{intentional} intentional overrides, {len(stale)} stale overrides"
    )
    return int(bool(different or unmatched or stale))


if __name__ == "__main__":
    raise SystemExit(main())
