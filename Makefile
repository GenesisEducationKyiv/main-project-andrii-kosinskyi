BINARY_NAME=app

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -gcflags="all=-N -l" -o build/app cmd/main.go

.PHONY: run
run:build
	 ./build/${BINARY_NAME}

.PHONY: format
format:
	@gofumpt -l -w .

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: lint
lint:
	@golangci-lint run ./... --config .golangci.yml
