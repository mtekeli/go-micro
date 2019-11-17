FROM golang:latest

RUN go get -u github.com/gorilla/mux

RUN go version

COPY . "/go/src/github/mtekeli/go-micro/app/backend"
WORKDIR /go/src/github/mtekeli/go-micro/app/backend

RUN go build -o primeNumberService

ENTRYPOINT [ "./primeNumberService" ]