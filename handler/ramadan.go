package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
)

func (h *Handler) RamadanHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = b.SendPhoto(ctx, &bot.SendPhotoParams{
		ChatID:      update.Message.Chat.ID,
		Photo:       &models.InputFileString{Data: "СЮДА_ВСТАВЬ_FILE_ID"},
		Caption:     messages.Messages[user.Language]["RamadanCaption"],
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(err)
		return
	}
}
