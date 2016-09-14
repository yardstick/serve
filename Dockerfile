FROM alpine
MAINTAINER Daniel Huckstep <danielh@getyardstick.com>

ADD serve /bin/

RUN mkdir /webroot
WORKDIR /webroot

EXPOSE 80

ENTRYPOINT ["serve"]
