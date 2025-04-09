# Этап 1: Сборка Go-приложения
FROM golang:1.24.2 AS builder

WORKDIR /home/namazbot
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Этап 2: Используем готовый образ с Chrome
FROM zenika/alpine-chrome:117-with-puppeteer

WORKDIR /home/namazbot

# Временно переключаемся на root
USER root

# Устанавливаем переменные окружения и часовой пояс
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/ \
    TZ="Asia/Dushanbe"

RUN apk add --no-cache tzdata

# Копируем собранное приложение
COPY --from=builder /home/namazbot .

# Делаем исполняемым
RUN chmod +x ./main

CMD ["./main"]
