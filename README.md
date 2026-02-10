# 🌳 tree-sitter-metasp [![ci-badge]][ci] [![pypi-version-badge]][pypi] [![npm-version-badge]][npm] [![crates-version-badge]][crates] [![rust-doc-badge]][rust-doc]

This repository provides the [tree-sitter] grammar for [metasp] language, a
system for implementing Answer Set Programming (ASP) systems via metaprogramming,
developed by the [Potassco][potassco] group.

## 📦 Installation

- **Python:** `pip install tree-sitter-metasp` ([PyPI][pypi])
- **Node.js:** `npm install tree-sitter-metasp` ([npm][npm])
- **Rust:** `cargo add tree-sitter-metasp` ([crates.io][crates])
- **C:** Build with `tree-sitter build`

## 📋 Release Checklist

We bundle generated files for easier deployment. Ensure to generate (build and
test) the parser using the following commands:

```bash
npx tree-sitter generate
npx tree-sitter build
npx tree-sitter test
```

When preparing a new release, ensure the version is updated consistently in the
following files:

- `package.json`
- `package-lock.json` (run `npm update -S` to update)
- `Cargo.toml`
- `Cargo.lock` (run `cargo update` to update)
- `pyproject.toml`
- `Makefile`

[tree-sitter]: https://tree-sitter.github.io/tree-sitter/
[metasp]: https://github.com/potassco/metasp
[potassco]: https://potassco.org/
[metasp-syntax.nvim]: https://github.com/rkaminsk/metasp-syntax.nvim
[ci-badge]: https://github.com/potassco/tree-sitter-metasp/workflows/CI%20test/badge.svg
[ci]: https://github.com/potassco/tree-sitter-metasp/actions/workflows/ci-test.yml
[crates-version-badge]: https://img.shields.io/crates/v/tree-sitter-metasp.svg
[crates]: https://crates.io/crates/tree-sitter-metasp
[rust-doc-badge]: https://img.shields.io/badge/api-rustdoc-blue.svg
[rust-doc]: https://docs.rs/tree-sitter-metasp
[pypi-version-badge]: https://img.shields.io/pypi/v/tree-sitter-metasp.svg
[pypi]: https://pypi.org/project/tree-sitter-metasp/
[npm-version-badge]: https://img.shields.io/npm/v/@potassco/tree-sitter-metasp.svg
[npm]: https://www.npmjs.com/package/@potassco/tree-sitter-metasp
