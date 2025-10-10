package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) AuthHandler(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message.Chat.ID != h.adminChatID {
			log.Println("unauthorized client")
			return
		}
		next(ctx, bot, update)
	}
}
