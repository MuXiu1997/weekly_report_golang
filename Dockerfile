FROM golang:latest as builder

COPY . /tmp/app

WORKDIR /tmp/app

RUN export GOPROXY=https://mirrors.aliyun.com/goproxy/ && \
go build -o ./main ./main.go && \
chmod +x ./main


FROM alpine:latest as prod

RUN groupadd -r app && useradd -r -g app app && mkdir -p /app

USER app

COPY --from=builder  --chown=app:app /tmp/app/main /tmp/app/template.docx /tmp/app/template/ /app/

EXPOSE 80

ENTRYPOINT ./main

