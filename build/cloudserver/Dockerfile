FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy \
    && go install github.com/swaggo/swag/cmd/swag@latest \
    && swag init \
    && go build -o cloudserver .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cloudserver .
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./cloudserver"]
