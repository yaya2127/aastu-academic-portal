# Multi-stage Docker build for Go Microservice API & React App
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o aastu-portal-server main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/aastu-portal-server .
COPY --from=builder /app/public ./public

EXPOSE 8080
CMD ["./aastu-portal-server"]
