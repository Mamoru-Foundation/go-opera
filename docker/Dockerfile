#FROM golang:1.17-alpine as builder
FROM golang:1.21 as builder

#RUN apk add --no-cache make gcc musl-dev linux-headers git
RUN apt-get update  \
    && apt-get install -y gcc musl-dev git curl tar libc6-dev \
    && apt-get clean
WORKDIR /go/go-opera
COPY . .

ARG GOPROXY
RUN go mod download
RUN make opera



#FROM alpine:latest
FROM debian:12.0-slim

RUN touch /var/run/supervisor.sock
#RUN apk add --no-cache ca-certificates bash supervisor
RUN apt-get update  \
    && apt-get install -y ca-certificates jq unzip bash grep curl sed htop procps cron  supervisor \
    && apt-get clean

COPY --from=builder /go/go-opera/build/opera /usr/bin/

EXPOSE 5050 18545 18546

ENTRYPOINT ["opera"]
