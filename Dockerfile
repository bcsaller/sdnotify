FROM golang:1.13-buster AS builder
WORKDIR /go/src/sdnotify
COPY . .
RUN go get -d -v ./...
RUN go install -buildmode=pie -v ./...

FROM ubuntu:latest
COPY --from=builder /go/bin/sdnotify /sdnotify
WORKDIR /
ENTRYPOINT [ "./sdnotify" ]
