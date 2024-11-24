FROM golang:1.23

WORKDIR backend

COPY . .

RUN go mod tidy

CMD ["go", "run", "main.go"]
