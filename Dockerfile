FROM golang:1.8-alpine

MAINTAINER Weerayut Hongsa <kusumoto.com@gmail.com>

COPY . /home/seedbox
WORKDIR /home/seedbox

RUN apk add --no-cache \ 
    git \
    build-base \
    curl \
    make \
    && go get gopkg.in/kataras/iris.v6 \
    && go get github.com/fsouza/go-dockerclient \
    && go build 

CMD ["bash","start.sh"]





