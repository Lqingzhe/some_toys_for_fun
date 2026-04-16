FROM golang:1.26-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum config.yml ./
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .

RUN go build -o App ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/App .
COPY --from=builder /app/config.yml .

CMD ["./App"]