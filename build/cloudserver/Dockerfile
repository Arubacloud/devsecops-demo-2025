FROM golang:1.21

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest

EXPOSE 8080

CMD swag init -g cmd/cloudserver/main.go && go run cmd/cloudserver/main.go