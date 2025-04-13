package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func (h *Handler) MailingMeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message.ReplyToMessage == nil {
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Ответьте на сообщение которое хотите отправить!",
			ReplyMarkup: inlineButtonMain(user.Language),
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

	_, err = b.ForwardMessage(ctx, &bot.ForwardMessageParams{
		ChatID:     update.Message.Chat.ID,
		FromChatID: update.Message.Chat.ID,
		MessageID:  update.Message.ReplyToMessage.ID,
	})
	if err != nil {
		log.Println(err)
		return
	}
}
