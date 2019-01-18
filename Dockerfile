FROM golang:1.11.4-stretch

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD go run main.go
