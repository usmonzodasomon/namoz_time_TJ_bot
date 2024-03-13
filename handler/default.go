package handler

import (
	"context"
	"echobot/messages"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func (h *Handler) DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
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
