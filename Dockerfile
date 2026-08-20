# syntax=docker/dockerfile:1

FROM golang:1.26@sha256:45a5f7a810238aabcbad211d70b9ae082022d96f7c7259e94041ad1b933575ac AS builder

WORKDIR /home/namazbot

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.19.1

WORKDIR /home/namazbot

RUN apk add --no-cache \
    tzdata \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    gzip \
    && apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main postgresql-client \
    && update-ca-certificates

ENV TZ="Asia/Dushanbe"
ENV ROD_BROWSER_PATH=/usr/bin/chromium-browser

COPY --from=builder /home/namazbot .

RUN chmod +x ./main

CMD ["./main"]
