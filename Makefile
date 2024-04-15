.PHONY: run local
local:
	@echo "Running local server"
	@cd cmd && go run main.go

.PHONY: generate schema graphql
generate:
	@echo "Generating schema"
	@go get -d github.com/99designs/gqlgen
	@go run github.com/99designs/gqlgen generate
