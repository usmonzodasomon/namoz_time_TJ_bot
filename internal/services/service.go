package services

import (
	"context"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/internal/domain"
)

type Service struct {
	storage   domain.Storage
	bot       domain.TelegramBot
	scheduler domain.Scheduler
}

func NewService(storage domain.Storage, bot domain.TelegramBot, scheduler domain.Scheduler) *Service {
	return &Service{
		storage:   storage,
		bot:       bot,
		scheduler: scheduler,
	}
}

func (s *Service) Start(ctx context.Context) error {
	// Запуск бота
	if err := s.bot.Start(ctx); err != nil {
		return err
	}

	// Запуск планировщика
	if err := s.scheduler.Start(ctx); err != nil {
		return err
	}

	// Здесь можно добавить другие задачи инициализации

	return nil
}

func (s *Service) Stop() error {
	// Остановка бота
	if err := s.bot.Stop(); err != nil {
		return err
	}

	// Остановка планировщика
	if err := s.scheduler.Stop(); err != nil {
		return err
	}

	return nil
}

// Добавьте здесь методы для обработки бизнес-логики
// Например:
func (s *Service) HandlePrayerTime(ctx context.Context, city string) error {
	// Логика обработки времени молитвы
	return nil
}

func (s *Service) SchedulePrayerNotifications(ctx context.Context) error {
	// Логика планирования уведомлений о молитве
	return nil
} 