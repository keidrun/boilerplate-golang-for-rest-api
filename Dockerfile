FROM golang:1.11.4-stretch AS builder
LABEL maintainer="Keid"
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:3.8
ENV APP_ENV=prod
WORKDIR /app
EXPOSE 3000
COPY --from=builder /app/server /app/
ENTRYPOINT ["/app/server"]
