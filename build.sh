#!/usr/bin/env bash
set -e
rm -f serve
docker run --rm -it -v "$PWD":/usr/src/app -w /usr/src/app golang:alpine ./compile.sh
tag=`git rev-parse --short HEAD`
docker build -t yardstick/serve:latest -t yardstick/serve:$tag .
docker push yardstick/serve:latest
docker push yardstick/serve:$tag
