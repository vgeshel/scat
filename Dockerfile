FROM golang
MAINTAINER vadim@yummly.com

COPY . /usr/src/app
WORKDIR /usr/src/app
RUN mkdir -p bin && CGO_ENABLED=0 go build -a -tags netgo -ldflags '-s' -o bin/scat scat.go
