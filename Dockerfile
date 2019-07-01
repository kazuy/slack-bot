FROM golang:1.12.6-alpine

RUN apk update && \
    apk add --update-cache --no-cache git && \
    rm -rf /var/cache/apk/*

ENV HOME /go/src/slack-bot
WORKDIR $HOME

COPY . $HOME
ENV GO111MODULE=on

CMD ["go", "run", "main.go"]
