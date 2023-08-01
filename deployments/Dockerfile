
FROM golang:1.20.5-alpine as build-stage
RUN apk --no-cache add ca-certificates
WORKDIR /go/src
COPY ./go.mod ./go.mod
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./cmd/app ./cmd/main.go

#
# final build stage
#
FROM scratch
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /go/src/cmd/app ./app
COPY --from=build-stage /go/src/_env ./_env
EXPOSE 8080

CMD ["./app"]