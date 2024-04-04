FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /home/app

COPY go.mod go.sum main.go /home/app

RUN go build

FROM alpine:3.19 AS prod

WORKDIR /home/app

COPY --from=build /home/app/Order /home/app

CMD ./Order
