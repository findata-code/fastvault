FROM golang:1.12.5-alpine3.9 as builder

RUN apk add --no-cache git

WORKDIR /builder
COPY . .

ENV GO111MODULE=on

RUN go mod download
RUN go build -o fastvault main.go

FROM alpine:3.9

RUN apk add --no-cache bash

WORKDIR /fastvault

COPY --from=builder /builder/fastvault .
COPY --from=builder /builder/resources/config.production.json .

ENV ENV_CONFIG_LOCATION=/fastvault/config.production.json

EXPOSE 8080

CMD ["/fastvault/fastvault"]