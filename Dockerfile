FROM golang:1.12.6-alpine

RUN apk update && \
    apk add --update-cache --no-cache git && \
    rm -rf /var/cache/apk/*

ENV HOME /go/src/github.com/kazuy/slack-bot
WORKDIR $HOME

COPY . $HOME
ENV GO111MODULE=on

EXPOSE 8080
CMD ["go", "run", "main.go"]
