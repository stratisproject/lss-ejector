FROM ubuntu:jammy

COPY build/lss-ejector /usr/local/bin

RUN apt-get update && apt-get -y install tzdata

ENTRYPOINT ["lss-ejector"]