REDOCLY_VERSION := latest
OAPI_CODEGEN_VERSION := latest

openapi:
	@echo "Generating code from OpenAPI specifications..."
	mkdir -p dist
	docker run --user=$(shell id -u):$(shell id -g) --rm -v$(PWD):/app -w /app redocly/cli:$(REDOCLY_VERSION) bundle spec/openapi/openapi.yml -o dist/openapi.yml
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@$(OAPI_CODEGEN_VERSION)
	oapi-codegen -generate types,client,server,spec,skip-prune -package generated -o spec/generated.go dist/openapi.yml
