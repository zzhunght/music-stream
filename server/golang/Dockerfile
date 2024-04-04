FROM golang:1.22

WORKDIR /app

COPY . .

COPY app.env app.env

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "cmd/app/main.go"]