FROM golang:1.12.12-alpine3.10 as builder

WORKDIR /usr/src/app

RUN apk add --no-cache git

COPY server server
COPY domain domain
COPY mod-3 mod-3

RUN cd mod-3 && go build

FROM alpine:3.10

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/mod-3/mod-3 .

ENTRYPOINT ["./mod-3"]
