# Этап 1: Собираем приложение Go
FROM golang:1.24.2 AS builder
WORKDIR /home/namazbot
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Этап 2: Используем образ с headless-shell
FROM chromedp/headless-shell:latest
WORKDIR /home/namazbot
COPY --from=builder /home/namazbot .

# Устанавливаем пакеты, включая tzdata
RUN apt-get update && apt-get install -y tzdata

ENV TZ="Asia/Dushanbe"

# Права на выполнение
RUN chmod +x ./main

ENTRYPOINT ["./main"]
