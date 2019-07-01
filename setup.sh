#!/bin/sh
docker build -t slack-bot .
docker run --rm -itd -v $PWD:/go/src/slack-bot -p 8080:8080 --name slack-bot slack-bot

