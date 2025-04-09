# Stage 1: Сборка Go-приложения
FROM golang:1.21 AS builder

WORKDIR /app

# Копируем go.mod и go.sum отдельно для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код и собираем
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Финальный образ с Chrome
FROM zenika/alpine-chrome:with-puppeteer

WORKDIR /app

# Устанавливаем переменные окружения
ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/ \
    TZ=Asia/Dushanbe

# Копируем собранное приложение
COPY --from=builder /app/main .

# Разрешаем запуск
RUN chmod +x ./main

# Запускаем
CMD ["./main"]
