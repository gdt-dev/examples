# Go Declarative Testing - Examples

[![Go Reference](https://pkg.go.dev/badge/github.com/gdt-dev/examples.svg)](https://pkg.go.dev/github.com/gdt-dev/examples)
[![Go Report Card](https://goreportcard.com/badge/github.com/gdt-dev/examples)](https://goreportcard.com/report/github.com/gdt-dev/examples)
[![Build Status](https://github.com/gdt-dev/examples/actions/workflows/gate-tests.yml/badge.svg?branch=main)](https://github.com/gdt-dev/examples/actions)
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg)](CODE_OF_CONDUCT.md)

<div style="float: left">
<img align=left src="static/gdtlogo400x544.png" width=200px />
</div>

`gdt` is a testing library that allows test authors to cleanly describe tests
in a YAML file. `gdt` reads YAML files that describe a test's assertions and
then builds a set of Golang structures that the standard Golang
[`testing`](https://golang.org/pkg/testing/) package and standard `go test`
tool can execute.

This `github.com/gdt-dev/examples` (`gdt-examples` hereafter) repository is a
companion Go library for `gdt` that contains examples for how to use `gdt` and
its companion Go libraries like `gdt-http` and `gdt-kube`.

The `http` directory in this repository contains Go code for a simplistic HTTP
web service that lists and creates books (in `http/api`) along with functional
tests written in YAML using `gdt`'s test file format (in `http/tests/api`).

## Contributing and acknowledgements

`gdt` was inspired by [Gabbi](https://github.com/cdent/gabbi), the excellent
Python declarative testing framework. `gdt` tries to bring the same clear,
concise test definitions to the world of Go functional testing.

The Go gopher logo, from which gdt's logo was derived, was created by Renee
French.

Contributions to `gdt-examples` are welcomed! Feel free to open a Github issue
or submit a pull request.
