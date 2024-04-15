FROM golang:alpine3.19 as go-builder

WORKDIR /app

COPY . .

ENV GONOPROXY="github.com/golang-jwt/jwt"

RUN apk add -u -t build-tools curl git && \
    apk add --no-cache ca-certificates && update-ca-certificates && \
    go build -o server *.go && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*

FROM alpine:latest

WORKDIR /app

COPY --from=go-builder /app/server /app/server

EXPOSE 8080

ENTRYPOINT [ "/app/server" ]
