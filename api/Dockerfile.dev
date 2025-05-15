FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod  ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]