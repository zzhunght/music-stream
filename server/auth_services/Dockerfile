FROM golang:1.22

WORKDIR /app

COPY . .

COPY app.env app.env

RUN go mod tidy

EXPOSE 30001

CMD ["go", "run", "main.go"]