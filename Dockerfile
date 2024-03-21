FROM golang:1.20.3-alpine AS builder

COPY . /github.com/Lokion13/auth-chat-server/source/
WORKDIR /github.com/Lokion13/auth-chat-server/source/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Lokion13/auth-chat-server/source/bin/auth_server .

CMD ["./auth_server"]