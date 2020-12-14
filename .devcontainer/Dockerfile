FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/yeyisan/DevContainer/main
RUN cd /build && git clone https://github.com/yeyisan/DevContainer.git

RUN cd /build/DevContainer/main && go build

EXPOSE 8080