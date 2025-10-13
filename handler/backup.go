package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) BackupHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("Manual backup requested by admin (chat ID: %d)", update.Message.Chat.ID)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "🔄 Создание резервной копии базы данных...",
	})
	if err != nil {
		log.Printf("Failed to send backup start message: %v", err)
	}

	if h.scheduler != nil {
		go h.scheduler.SendDatabaseBackup()
	} else {
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "❌ Ошибка: планировщик не инициализирован",
		})
		if err != nil {
			log.Printf("Failed to send error message: %v", err)
		}
	}
}
