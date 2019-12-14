FROM golang:latest

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/apsdehal/go-logger

RUN go version

COPY . "/go/src/github.com/mtekeli/go-micro/app/frontend"
WORKDIR /go/src/github.com/mtekeli/go-micro/app/frontend

RUN go build -o webApp

ENTRYPOINT [ "./webApp" ]