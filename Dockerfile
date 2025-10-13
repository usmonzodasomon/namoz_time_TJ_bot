FROM golang:1.24.2 AS builder

WORKDIR /home/namazbot
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.19.1

WORKDIR /home/namazbot
COPY --from=builder /home/namazbot .

ENV TZ="Asia/Dushanbe"
ENV ROD_BROWSER_PATH=/usr/bin/chromium-browser

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

RUN chmod +x ./main

CMD ["./main"]
