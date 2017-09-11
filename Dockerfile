FROM alpine:latest

RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*

RUN apk add --update go git

ADD caddy /caddy

ADD Caddyfile /Caddyfile

ADD readme.md /readme.md

ENTRYPOINT ["/caddy"]

