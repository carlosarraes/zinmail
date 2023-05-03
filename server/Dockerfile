FROM golang:1.16-alpine

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY . .

RUN go build -o bin/main

EXPOSE 8080

CMD ["./bin/main"]
