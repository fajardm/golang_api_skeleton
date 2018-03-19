FROM golang

ADD . /go/src/github.com/fajardm/golang_api_skeleton

RUN go install github.com/fajardm/golang_api_skeleton

ENTRYPOINT /go/bin/golang_api_skeleton

EXPOSE 3000
