FROM golang:latest

ARG db_addr
ARG db_psw
ARG db_user
ARG db_name

ENV DB_ADDR=$db_addr
ENV DB_PSW=$db_psw
ENV DB_USER=$db_user
ENV DB_NAME=$db_name

WORKDIR /usr/src/app

COPY go.work ./
COPY libs/go ./libs/go/
COPY ./deploy/server ./deploy/server/

RUN go build -v -o /usr/local/bin/app ./deploy/server

CMD ["app"]