#!/bin/bash

# Конфигурация
CONTAINER_NAME="namozbot"
IMAGE_NAME="namozbot"
REPO_PATH="/home/somon/namoz_time_TJ_bot"

echo "🚀 Начинаем деплой..."

# Переходим в директорию проекта
cd $REPO_PATH

# Пуллим последние изменения
echo "📥 Получаем последние изменения..."
git pull origin main

# Останавливаем контейнер
echo "⏹️ Останавливаем контейнер..."
sudo docker stop $CONTAINER_NAME 2>/dev/null || echo "Контейнер уже остановлен"

# Удаляем контейнер
echo "🗑️ Удаляем контейнер..."
sudo docker rm $CONTAINER_NAME 2>/dev/null || echo "Контейнер уже удален"

# Удаляем образ
echo "🗑️ Удаляем старый образ..."
sudo docker rmi $IMAGE_NAME 2>/dev/null || echo "Образ уже удален"

# Создаем новый образ
echo "🔨 Создаем новый образ..."
sudo docker build -t $IMAGE_NAME .

# Запускаем контейнер
echo "🏃 Запускаем контейнер..."
sudo docker run --name namozbot \
  --restart on-failure:5 \
  -v /home/somon/namoz_time_TJ_bot/data/logs:/home/namazbot/data/logs \
  -d namozbot

echo "✅ Деплой завершен!"

# Показываем статус
sudo docker ps | grep $CONTAINER_NAME