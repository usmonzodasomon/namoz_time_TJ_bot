# Этап 1: Сборка Go-приложения
FROM golang:1.24.2 AS builder

WORKDIR /home/namazbot
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Этап 2: Используем готовый образ с Chrome
FROM zenika/alpine-chrome@sha256:9b3e4cb7a83f2f5e2c2176d29a2d631c693f42bfc173e13fd00578c83cf99bbf

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
