FROM ubuntu:latest
LABEL maintainer="inclunet"
WORKDIR /pkg
COPY ./reports/templates/*.* /accessbot/reports/templates/*.*
COPY ./accessbot /usr/bin/accessbot
CMD accessbot
