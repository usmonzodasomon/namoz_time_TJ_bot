package domain

import (
	"context"
	"github.com/go-telegram/bot"
)

// Storage интерфейс для работы с хранилищем данных
type Storage interface {
	// Добавьте методы для работы с данными
}

// TelegramBot интерфейс для работы с Telegram
type TelegramBot interface {
	Start(ctx context.Context) error
	Stop() error
	SendMessage(chatID int64, text string) error
}

// Scheduler интерфейс для планировщика задач
type Scheduler interface {
	Start(ctx context.Context) error
	Stop() error
	ScheduleTask(task func()) error
}

// Handler интерфейс для обработки сообщений
type Handler interface {
	HandleMessage(ctx context.Context, msg *bot.Message) error
} 