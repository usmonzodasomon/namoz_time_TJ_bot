package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/internal/config"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/internal/infrastructure/database"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/internal/infrastructure/telegram"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/internal/infrastructure/scheduler"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/internal/services"
)

func main() {
	// Инициализация логгера
	logger := initLogger()
	logger.Println("Starting application...")

	// Загрузка конфигурации
	if err := godotenv.Load(); err != nil {
		logger.Printf("Error loading .env file: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	}

	// Создание контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Инициализация компонентов
	db, err := database.NewDatabase(cfg.Database)
	if err != nil {
		logger.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Инициализация Telegram бота
	bot, err := telegram.NewBot(cfg.Telegram.Token)
	if err != nil {
		logger.Fatalf("Failed to initialize Telegram bot: %v", err)
	}

	// Инициализация планировщика
	scheduler := scheduler.NewScheduler()

	// Инициализация сервисов
	service := services.NewService(db, bot, scheduler)

	// Запуск сервиса
	go func() {
		if err := service.Start(ctx); err != nil {
			logger.Printf("Service error: %v", err)
			cancel()
		}
	}()

	// Обработка сигналов завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		logger.Printf("Received signal: %v", sig)
		cancel()
	case <-ctx.Done():
		logger.Println("Context cancelled")
	}

	// Graceful shutdown
	if err := service.Stop(); err != nil {
		logger.Printf("Error during shutdown: %v", err)
	}
}

func initLogger() *log.Logger {
	file, err := os.OpenFile("data/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", log.LstdFlags)
} 