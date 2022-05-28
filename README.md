# FastlyWasmly

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/grokify/fastlywasmly/workflows/test/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/fastlywasmly/actions/workflows/go_build.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/fastlywasmly
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/fastlywasmly
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/fastlywasmly
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/fastlywasmly
 [license-svg]: https://img.shields.io/badge/license-MIT-fastlywasmly.svg
 [license-url]: https://github.com/grokify/fastlywasmly/blob/master/LICENSE

`fastlywasmly` is a CLI app that will package a Fastly Compute@Edge TOML and WASM file into a Fastly tarball. It can alternately take a `bin` directly in place of the WASM file to load all the files in the directory and subdirectories.

It builds the tarball filename and internal folder structure using the same same approach as Fastly CLI.

## Installation

```bash
% go install github.com/grokify/fastlywasmly
```

## Usage

Loading with WASM file:

```bash
% fastlywasmly -t /path/to/fastly.toml -w /path/to/yourname.wasm
```

Loading with bin dir (including WASM file):

```bash
% fastlywasmly -5 /path/to/fastly.toml -b /path/to/bin
```