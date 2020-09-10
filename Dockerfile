FROM golang:1.12-alpine

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN apk update && \
    apk add --no-cache git make npm gcc sudo

RUN go get -u github.com/pilu/fresh
RUN npm init --yes
RUN npm install -g yarn webpack webpack-cli
RUN npm i -g create-react-app

# docker-compose up 前に実行
# docker-compose run --rm app sh -c "create-react-app client && cd client && yarn add webpack-cli react react-dom react-redux redux react-router-dom redux-persist axios"

# RUN go build -o bin/app main.go
# CMD ["fresh"]
