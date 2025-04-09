FROM golang:1.24.2 AS builder
WORKDIR /home/namazbot
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.19.1
WORKDIR /home/namazbot

# Установка Chromium и зависимостей
RUN apk update && apk add --no-cache \
    chromium=118.0.5993.117-r0 \
    chromium-chromedriver=118.0.5993.117-r0 \
    harfbuzz \
    nss \
    freetype \
    ttf-freefont \
    font-noto-emoji \
    wqy-zenhei \
    tzdata

# Настройка переменных окружения
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/ \
    TZ="Asia/Dushanbe"

# Копирование собранного приложения
COPY --from=builder /home/namazbot .

# Разрешения на выполнение
RUN chmod +x ./main

# Запуск приложения
CMD ["./main"]