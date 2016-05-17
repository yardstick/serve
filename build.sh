#!/usr/bin/env bash
docker run --rm -it -v "$PWD":/usr/src/app -w /usr/src/app golang:alpine go build serve.go
docker build -t darkhelmetlive/serve .
docker push darkhelmetlive/serve
