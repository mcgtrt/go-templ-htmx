build:
	@go build -o ./bin/api

run: templ build
	@./bin/api

test:
	@go test -v ./... -count=1

templ:
	@./templ generate

.PHONY: templ