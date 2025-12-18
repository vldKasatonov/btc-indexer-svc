FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/vldKasatonov/btc-indexer-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/btc-indexer-svc /go/src/github.com/vldKasatonov/btc-indexer-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/btc-indexer-svc /usr/local/bin/btc-indexer-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["btc-indexer-svc"]
