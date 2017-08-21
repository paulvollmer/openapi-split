build:
	@go build

test: build
	./openapi-split -f json -i fixtures/index.yaml -d fixtures/definitions -p fixtures/paths -r fixtures/responses

release:
	goreleaser --skip-validate --rm-dist --skip-publish
