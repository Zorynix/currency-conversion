FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /app/main ./cmd/app/main.go

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow

ENV TZ Europe/Moscow

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /build/config/config.yaml /app/config/config.yaml
COPY --from=builder /build/migrations /app/migrations

CMD ["/app/main"]
