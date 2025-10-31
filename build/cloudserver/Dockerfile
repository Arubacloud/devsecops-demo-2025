FROM golang:1.21

WORKDIR /app
COPY . .

# Use direct mode
ENV GOPROXY=https://goproxy.io,direct
ENV GOPRIVATE=""
ENV GOSUMDB=off

RUN go mod download
RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest

EXPOSE 8080

CMD swag init -g cmd/restapi/main.go && go run cmd/restapi/main.go