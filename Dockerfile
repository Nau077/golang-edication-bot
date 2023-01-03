FROM golang:1.17.5-alpine3.15 AS builder

COPY . /go-edication-bot/
WORKDIR /go-edication-bot/

RUN go mod download
RUN go build -o ./bin/go-edication-bot/ ./cmd/main.go ./static

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go-edication-bot/bin/go-edication-bot .
COPY --from=builder /go-edication-bot/static/ static/

EXPOSE 80

CMD ["./go-edication-bot"]