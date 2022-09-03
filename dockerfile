FROM golang:1.18.5 AS build

WORKDIR /fddns

COPY . /fddns/

ENV GOPROXY https://goproxy.cn,direct

RUN export GOPROXY=https://goproxy.cn,direct

RUN go build -tags netgo -o ./fddns

# FROM ubuntu:22.04
# RUN apt-get -qq update \
#     && apt-get -qq install -y --no-install-recommends ca-certificates curl

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /fddns/fddns /

CMD /fddns