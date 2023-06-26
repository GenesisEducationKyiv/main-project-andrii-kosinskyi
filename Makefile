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

.PHONY: rate
rate:
	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:8080/api/rate

.PHONY: subscribe
subscribe:
	curl -X POST -F 'email=kosinskiy.andrey@ukr.net' http://localhost:8080/api/subscribe

.PHONY: subscribe2
subscribe2:
	curl -X POST -F 'email=andrey.kosinskiy@hellotickets.com' http://localhost:8080/api/subscribe

.PHONY: subscribeInvalid
subscribeInvalid:
	curl -X POST -F 'email=andrey.kosinskiy' http://localhost:8080/api/subscribe

.PHONY: sendMails
sendMails:
	curl -X POST http://localhost:8080/api/sendEmails