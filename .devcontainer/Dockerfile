FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/yeyisan/DevContainer
RUN cd /build && git clone https://github.com/yeyisan/DevContainer.git

RUN cd /build/DevContainer && go build -o devcontainer

EXPOSE 8080

ENTRYPOINT ["/build/DevContainer/devcontainer"]