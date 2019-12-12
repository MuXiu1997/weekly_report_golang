FROM golang:1.13-alpine3.10 as builder

COPY . /tmp/app

WORKDIR /tmp/app

RUN export GOPROXY=https://mirrors.aliyun.com/goproxy/ && \
go build -o ./main ./main.go && \
chmod +x ./main


FROM alpine:latest as prod

RUN set -x \
    && addgroup -g 1000 -S app \
    && adduser -S -G app -u 999 app \
    && mkdir -p /app && chown app:app /app \
    && mkdir -p /app/data && chown app:app /app/data

USER app

WORKDIR /app

VOLUME /app/data

COPY --from=0  --chown=app:app /tmp/app/main /tmp/app/template.docx ./

COPY --from=0  --chown=app:app /tmp/app/template/ ./template/

EXPOSE 8080

CMD ["./main"]
