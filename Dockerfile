# Используем Go на базе Alpine
FROM golang:1.24-alpine AS builder

WORKDIR /home/namazbot
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Финальный контейнер
FROM alpine:3.19.1

WORKDIR /home/namazbot
COPY --from=builder /home/namazbot .

# Настройки окружения
ENV TZ="Asia/Dushanbe"
ENV ROD_BROWSER_PATH=/usr/bin/chromium-browser

# Установка зависимостей и Chromium
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    && update-ca-certificates \
    # Убираем лишнее, чтобы уменьшить размер
    && rm -rf /var/cache/apk/*

RUN chmod +x ./main

CMD ["./main"]
