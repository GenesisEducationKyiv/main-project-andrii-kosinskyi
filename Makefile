.PHONY: run
run:
	@go run cmd/main.go

.PHONY: format
format:
	@gofumpt -l -w .

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: lint
lint:
	@golangci-lint run ./... --config .golangci.yml
