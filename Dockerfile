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
EXPOSE 8080

CMD ["sleep", "10000"]