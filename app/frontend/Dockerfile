FROM golang:latest

RUN go get -u github.com/gorilla/mux

RUN go version

COPY . "/go/src/github/mtekeli/go-micro/app/frontend"
WORKDIR /go/src/github/mtekeli/go-micro/app/frontend

RUN go build -o webApp

ENTRYPOINT [ "./webApp" ]