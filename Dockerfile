FROM golang:1.26

WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
