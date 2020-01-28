FROM golang:1.13.6-alpine

ENV TERM screen-256color

RUN apk add --update \
    curl \
    git \
    unzip \
    libzip-dev \
    libgcc \
    libbz2 \
    bzip2-dev \
    autoconf \
    make \
    zlib \
    zlib-dev \
    g++

RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

RUN mkdir -p /var/apps/cuvva

VOLUME /var/apps/cuvva

WORKDIR /var/apps/cuvva