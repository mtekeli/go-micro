FROM golang:latest

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/apsdehal/go-logger

RUN go version

COPY . "/go/src/github.com/mtekeli/go-micro/app/backend"
WORKDIR /go/src/github.com/mtekeli/go-micro/app/backend

RUN go build -o primeNumberService

ENTRYPOINT [ "./primeNumberService" ]