FROM golang:1.12-alpine

ENV GO111MODULE=on

WORKDIR /go/src/app/golang
COPY . .

RUN apk update && \
    apk add --no-cache git make gcc sudo

RUN go get github.com/pilu/fresh

# RUN go build -o bin/app main.go
CMD ["fresh"]
