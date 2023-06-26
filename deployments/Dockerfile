
FROM golang:1.20.5-alpine as build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/bitcoin_checker_api

COPY .. .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /cmd


#
# final build stage
#
FROM scratch

# Copy ca-certs for app web access
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /bitcoin_checker_api /bitcoin_checker_api
WORKDIR /go/src/bitcoin_checker_api
# app uses port 5000
EXPOSE 5000

ENTRYPOINT ["/bitcoin_checker_api/build/app"]