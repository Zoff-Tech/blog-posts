
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o send-email

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/send-email .
EXPOSE 8080
CMD ["./send-email"]
