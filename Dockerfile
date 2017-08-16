FROM alpine:latest

RUN apk --update add ca-certificates && rm -rf /var/cache/apk/*

RUN apk add --update go git

ADD bin/caddy /caddy

ADD Caddyfile /Caddyfile

ADD readme.md /readme.md

ENTRYPOINT ["/caddy"]



# docker build -t mickeyyawn/scribl:latest .
# docker run -p 127.0.0.1:9001:8080 --env GITHUB_URL=github.com/mickeyyawn/yawn.me -i mickeyyawn/scribl:latest
# docker run -p 127.0.0.1:9001:8080 --env GITHUB_URL=github.com/turnerlabs/turnerlabs.io -i mickeyyawn/scribl:latest