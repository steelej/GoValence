#!/usr/bin/env python3
"""Download JSON-like model definitions throughout the Brightspace docs.

Install Beautiful Soup with ``python -m pip install -r docs/models/requirements.txt``.
Run from any directory with ``python docs/models/scrape_models.py``. Use ``--page``
to scrape one page while checking the output format. The default crawl uses the
site sitemap, including resource pages, basic API pages, and glossary entries.
"""

from __future__ import annotations

import argparse
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path
import re
import sys
from urllib.error import URLError
from urllib.parse import urljoin, urlparse
from urllib.request import Request, urlopen
from xml.etree import ElementTree

from bs4 import BeautifulSoup


INDEX_URL = "https://docs.valence.desire2learn.com/sitemap.xml"
OUTPUT_DIR = Path(__file__).resolve().parent
NAME_PART = re.compile(r"^[A-Za-z_][A-Za-z_0-9-]*$")


def download_bytes(url: str) -> bytes:
    request = Request(url, headers={"User-Agent": "GoValence-model-scraper/1.0"})
    with urlopen(request, timeout=30) as response:
        return response.read()


def download(url: str) -> BeautifulSoup:
    return BeautifulSoup(download_bytes(url), "html.parser")


def reference_pages(index_url: str) -> list[str]:
    data = download_bytes(index_url)
    if urlparse(index_url).path.endswith(".xml"):
        root = ElementTree.fromstring(data)
        links = [
            element.text
            for element in root.iter()
            if element.tag.rsplit("}", 1)[-1] == "loc"
        ]
    else:
        index = BeautifulSoup(data, "html.parser")
        links = [link["href"] for link in index.select("a[href]")]
    pages = set()
    for link in links:
        if not link:
            continue
        page = urljoin(index_url, link).split("#", 1)[0]
        parsed = urlparse(page)
        if parsed.netloc == urlparse(index_url).netloc and parsed.path.endswith(".html"):
            pages.add(page)
    if not pages:
        raise ValueError(f"No HTML pages found at {index_url}")
    return sorted(pages)


def code_blocks(body) -> list[str]:
    blocks = [pre.get_text().strip("\n") for pre in body.select("pre")]
    return [block for block in blocks if block.lstrip().startswith(("{", "["))]


def page_group(page: BeautifulSoup) -> str:
    heading = page.select_one("h2.heading")
    label = heading.get_text(" ", strip=True) if heading else "General"
    return "".join(word[:1].upper() + word[1:] for word in re.findall(r"[A-Za-z0-9]+", label))


def models_on_page(page_url: str) -> list[tuple[str, str, str]]:
    page = download(page_url)
    models = []
    default_group = page_group(page)
    for definition in page.select("dl.js.attribute"):
        heading = definition.find("dt", recursive=False)
        body = definition.find("dd", recursive=False)
        if heading is None or body is None:
            continue
        anchor = heading.get("id", "")
        name = anchor if "." in anchor else f"{default_group}.{anchor}"
        parts = name.split(".")
        if len(parts) < 2 or not all(NAME_PART.fullmatch(part) for part in parts):
            continue
        blocks = code_blocks(body)
        if blocks:
            description = "\n\n".join(blocks)
        else:
            prose = body.get_text(" ", strip=True)
            if not prose:
                continue
            description = f"[No JSON-like schema in source]\n{prose}"
        if not description:
            continue
        direct_url = f"{page_url}#{anchor}"
        models.append((name, direct_url, description))
    for definition in page.select("dl.glossary"):
        for heading in definition.find_all("dt", recursive=False):
            body = heading.find_next_sibling("dd")
            if body is None or body.find_previous_sibling("dt") is not heading:
                continue
            anchor = heading.get("id", "")
            term = anchor.removeprefix("term-")
            if not anchor.startswith("term-") or not NAME_PART.fullmatch(term):
                continue
            blocks = code_blocks(body)
            if blocks:
                models.append((f"Common.{term}", f"{page_url}#{anchor}", "\n\n".join(blocks)))
    return models


def scrape(pages: list[str], output_dir: Path) -> tuple[int, int]:
    # Map preserves the first reference URL if a model is repeated on another page.
    models: dict[str, tuple[str, str]] = {}
    with ThreadPoolExecutor(max_workers=4) as pool:
        for page_url, page_models in zip(pages, pool.map(models_on_page, pages)):
            print(f"{page_url}: {len(page_models)} models")
            for name, direct_url, description in page_models:
                previous = models.get(name)
                if previous is not None and previous[1] != description:
                    raise ValueError(f"Conflicting definitions for {name}: {previous[0]} and {direct_url}")
                models.setdefault(name, (direct_url, description))

    # Fetch and validate everything before changing files.
    for name, (direct_url, description) in sorted(models.items()):
        group, model = name.split(".", 1)
        destination = output_dir / group / model
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.write_text(f"{direct_url}\n\n{description}\n", encoding="utf-8")
    return len(models), len(pages)


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--page", help="Scrape only this resource page URL")
    parser.add_argument("--index-url", default=INDEX_URL)
    parser.add_argument("--output-dir", type=Path, default=OUTPUT_DIR)
    args = parser.parse_args()
    try:
        pages = [args.page] if args.page else reference_pages(args.index_url)
        count, page_count = scrape(pages, args.output_dir)
    except (URLError, OSError, ValueError) as exc:
        print(f"Scrape failed: {exc}", file=sys.stderr)
        return 1
    print(f"Wrote {count} models from {page_count} pages to {args.output_dir}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
