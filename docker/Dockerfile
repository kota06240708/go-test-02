FROM golang:1.12.0-alpine3.9

ENV SRC_DIR=/go/src/github.com/api

WORKDIR $SRC_DIR

ADD . $SRC_DIR

ENV GO111MODULE=on

RUN apk add --no-cache \
  alpine-sdk \
  git \
  && go get github.com/pilu/fresh

RUN go get github.com/jinzhu/gorm
RUN go get github.com/gin-gonic/gin

# RUN go get github.com/jinzhu/gorm \
#   go get bitbucket.org/liamstask/goose/cmd/goose \
#   go get github.com/gin-gonic/gin

EXPOSE 4000

CMD ["fresh"]