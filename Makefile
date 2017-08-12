build:
	@go build

test: build
	./openapi-split -i fixtures/index.yaml -d fixtures/definitions -p fixtures/paths -r fixtures/responses
