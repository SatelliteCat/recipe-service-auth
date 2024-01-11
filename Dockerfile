FROM golang:1.21.5-alpine3.19 AS builder

COPY . /auth/source/
WORKDIR /auth/source/

RUN go mod download
RUN go build -o ./bin/http_server cmd/http_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /auth/source/bin/http_server .
COPY --from=builder /auth/source/.env .

CMD sh -c 'ls -la'
CMD ["./http_server"]
