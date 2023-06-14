FROM ubuntu:latest AS builder
LABEL maintainer="inclunet"
COPY ./reports/templates/*.* /accessbot/reports/templates/*.*
COPY ./cmd/accessbot/accessbot /usr/bin/accessbot
CMD accessbot
