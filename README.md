# OpenAPI Split [![Build Status](https://travis-ci.org/paulvollmer/openapi-split.svg?branch=master)](https://travis-ci.org/paulvollmer/openapi-split)

`openapi-split` is a simple commandline tool to organize an openapi (swagger) specification across multiple files.
the tool read `yaml` files from different directories like `paths`, `definitions` and `responses` to build a single openapi spec file.

**THIS IS AN ALPHA RELEASE. KEEP IN MIND THERE CAN BE MAJOR CHANGES IN FUTURE**


## Installation
If you have installed `go` on your system you can run the following command to build and install the tool.
```
go get github.com/paulvollmer/openapi-split
```
Or you can download the latest binary from the [release github page](https://github.com/paulvollmer/openapi-split/releases).


## Usage
The philosophy behind the tool is to KISS (keep it stupid simple) at the usage.

```
openapi-split > spec.yaml

# or encode the data to json with the -f flag
openapi-split -f json > spec.json
```
