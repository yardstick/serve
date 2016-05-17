FROM alpine
MAINTAINER Daniel Huckstep <darkhelmet@darkhelmetlive.com>

ADD serve /bin/

RUN mkdir /html
WORKDIR /html

ENTRYPOINT ["serve"]
