FROM golang:1.12-alpine

WORKDIR /go/src/app
COPY . .

RUN apk update && \
    apk add --no-cache git npm sudo

RUN npm init --yes
RUN npm install -g yarn webpack webpack-cli
RUN npm i -g create-react-app