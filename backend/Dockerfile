FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 42069

CMD ["go", "run", "main.go"]
