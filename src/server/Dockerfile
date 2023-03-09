FROM golang:1.19-alpine as dev

WORKDIR /go/src/app

COPY ./go.mod go.mod
COPY ./go.sum go.sum
RUN go mod download
COPY . .
RUN go build

CMD ./mantras
