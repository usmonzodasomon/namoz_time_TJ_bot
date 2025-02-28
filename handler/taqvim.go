package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"echobot/messages"
)

func (h *Handler) TaqvimHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		user.Language = "tj"
	}

	_, err = b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:      update.Message.Chat.ID,
		Photo:       &models.InputFileString{Data: "AgACAgIAAxkBAAE-McFnwdCAhqF7jCNOEIMh1H3pssYl9AACc_ExGzDvEErEkzsnLELdoAEAAwIAA3kAAzYE"},
		Caption:     messages.Messages[user.Language]["Taqvim"],
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(fmt.Sprintf("taqvimhandler error: %v", err))
		return
	}
}
