FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY cmd/bot/main.go .
COPY . .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]




