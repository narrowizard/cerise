FROM golang as builder
WORKDIR /go/src/github.com/narrowizard/cerise
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -o cerise .

FROM alpine:latest

ENV TIMEZONE=Asia/Shanghai

VOLUME ["/root/config", "/root/log"]
WORKDIR /root/

RUN rm -rf /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/v3.5/community/" >> /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/v3.5/main/" >> /etc/apk/repositories && \
    apk add --update ca-certificates tzdata
RUN for cert in `ls -1 /etc/ssl/certs/*.crt | grep -v /etc/ssl/certs/ca-certificates.crt`; do cat "$cert" >> /etc/ssl/certs/ca-certificates.crt; done

COPY --from=builder /go/src/github.com/narrowizard/cerise/cerise .
COPY --from=builder /go/src/github.com/narrowizard/cerise/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /go/src/github.com/narrowizard/cerise/entrypoint.sh .

CMD ["/bin/sh", "-c", "./entrypoint.sh"]