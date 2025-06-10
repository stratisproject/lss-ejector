FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache make git gcc g++ musl-dev linux-headers

COPY go.mod /lss-ejector/
COPY go.sum /lss-ejector/
RUN cd /lss-ejector && go mod download
RUN go env -w CGO_ENABLED=1

ADD . /lss-ejector
RUN cd /lss-ejector && make build-linux

FROM alpine:latest

RUN apk add --no-cache ca-certificates libstdc++ tzdata
COPY --from=builder /lss-ejector/build/lss-ejector /usr/local/bin/

ENTRYPOINT ["lss-ejector"]