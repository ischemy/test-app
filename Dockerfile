FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o test-app .

FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/test-app .
EXPOSE 8080
CMD ["./test-app"]
