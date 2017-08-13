# OpenAPI Split [![Build Status](https://travis-ci.org/paulvollmer/openapi-split.svg?branch=master)](https://travis-ci.org/paulvollmer/openapi-split)

`openapi-split` is a simple commandline tool to organize an openapi (swagger) specification across multiple files.
the tool read yaml files from different directories like `paths`, `definitions` and `responses` to build a single openapi spec file.

**THIS IS AN ALPHA RELEASE. KEEP IN MIND THERE CAN BE MAJOR CHANGES IN FUTURE**


## Installation
```
go get github.com/paulvollmer/openapi-split
```


## Usage
The philosophy behind the tool is to KISS (keep it stupid simple) at the usage.

```
openapi-split > spec.yaml
```
