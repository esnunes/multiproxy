FROM golang:alpine AS build-env

RUN apk update && apk add git

RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/UnnoTed/fileb0x

ADD . /go/src/github.com/esnunes/multiproxy

RUN cd /go/src/github.com/esnunes/multiproxy/ && \
    dep ensure -v && \
    fileb0x assets.yaml && \
    go install ./...

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/bin/multiproxy /app/
ENTRYPOINT ["./multiproxy"]
CMD ["./config.json"]
