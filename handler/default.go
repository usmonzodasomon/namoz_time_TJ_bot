package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"log"
)

func (h *Handler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	if update.Message.Photo != nil && len(update.Message.Photo) > 0 {
		log.Println("file_id:", update.Message.Photo[len(update.Message.Photo)-1].FileID)
	}
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["UnknownCommand"],
		ReplyMarkup: inlineButtonMain(user.Language),
	})
}
