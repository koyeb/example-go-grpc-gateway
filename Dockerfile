FROM golang:1.19-alpine AS builder
ARG ORDER_SERVICE_ADDRESS
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./orders-service ./cmd/server/main.go
RUN go build -ldflags  "-X main.orderServiceAddr=$ORDER_SERVICE_ADDRESS" -o ./gateway-service ./cmd/client/main.go
 
FROM alpine:latest AS orders-service
WORKDIR /app
COPY --from=builder /app/orders-service .
EXPOSE 50051
ENTRYPOINT ["./orders-service"]

FROM alpine:latest AS gateway-service
WORKDIR /app
COPY --from=builder /app/gateway-service .
EXPOSE 8080
ENTRYPOINT ["./gateway-service"]
