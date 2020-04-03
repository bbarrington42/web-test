#!/usr/bin/env bash

set -ev

go test -v ./...

docker build -t web-test .

docker run --name web-test -p 127.0.0.1:80:8080 -d web-test

PONG=$(curl http://localhost/ping)

docker container stop web-test
docker container rm web-test

if [ "$PONG" = "pong" ]; then
  exit 0
else
  exit 1
fi
