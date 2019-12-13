# syntax = docker/dockerfile:experimental
FROM golang:1.13-alpine3.10 as builder

ENV GOSU_VERSION 1.11

COPY . /tmp/app

WORKDIR /tmp/app

RUN --mount=type=cache,target=/go,id=go_path,sharing=locked \
    export GOPROXY=https://mirrors.aliyun.com/goproxy/ \
    && go build -o ./main ./main.go \
    && chmod +x ./main

RUN --mount=type=cache,target=/tmp/gosu,id=gosu,sharing=locked \
    echo 'if ! [ -f /tmp/gosu/gosu ]; then' > /tmp/gosu/get_gosu.sh \
    && echo '  wget -O /tmp/gosu/gosu "https://github.com/tianon/gosu/releases/download/$GOSU_VERSION/gosu-amd64"' >> /tmp/gosu/get_gosu.sh \
    && echo 'fi' >> /tmp/gosu/get_gosu.sh \
    && chmod +x /tmp/gosu/get_gosu.sh \
    && /tmp/gosu/get_gosu.sh \

    && chmod +x /tmp/gosu/gosu \
    && /tmp/gosu/gosu nobody true


FROM alpine:latest as prod

RUN --mount=type=cache,target=/tmp/app,from=builder,source=/tmp/app \
    --mount=type=cache,target=/tmp/gosu,id=gosu,sharing=locked \
    set -x \
    && addgroup -g 1000 -S app \
    && adduser -S -G app -u 1000 app \
    && mkdir -p /app \
    && mkdir -p /app/data \

    && cp /tmp/app/main /app \
    && cp /tmp/app/template.docx /app \
    && cp -R /tmp/app/template/ /app/template/ \
    && chown -R app:app /app \

    && cp /tmp/gosu/gosu /usr/local/bin/gosu \

    && cp /tmp/app/docker-entrypoint.sh /usr/local/bin/ \
    && chmod +x /usr/local/bin/docker-entrypoint.sh

WORKDIR /app

VOLUME /app/data

ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 8080

CMD ["/app/main"]
