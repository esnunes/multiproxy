FROM golang:alpine AS build-env

RUN apk update && apk add git

RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/github.com/esnunes/multiproxy

RUN cd /go/src/github.com/esnunes/multiproxy/ && \
    dep ensure -v && \
    go build -o multiproxy cmd/multiproxy/main.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/esnunes/multiproxy/multiproxy /app/
ENTRYPOINT ./multiproxy
