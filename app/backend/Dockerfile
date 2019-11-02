FROM golang:latest

RUN go get -u github.com/gorilla/mux

WORKDIR /app

COPY . .

RUN go build -o primeNumberService

ENTRYPOINT [ "./primeNumberService" ]