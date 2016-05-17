#!/usr/bin/env bash
set -e
rm -f serve
docker run --rm -it -v "$PWD":/usr/src/app -w /usr/src/app golang:alpine ./compile.sh
docker build -t darkhelmetlive/serve .
docker push darkhelmetlive/serve
