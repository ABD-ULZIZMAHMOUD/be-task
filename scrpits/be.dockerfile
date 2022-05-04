FROM golang:1.17-buster as builder

WORKDIR /be-task

ENV GO111MODULE=on
COPY ./ .
RUN go get -tags musl ./...
RUN go mod download
RUN go mod vendor

CMD ["go","run","main.go"]


