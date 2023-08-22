FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]