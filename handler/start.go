package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"log"
)

func (h *Handler) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user := types.User{
		ChatID:   update.Message.Chat.ID,
		RegionID: 1,
		Username: update.Message.From.Username,
		Language: "tj",
	}

	if err := h.storage.AddUserIfNotExist(user); err != nil {
		log.Println(err)
		return
	}

	//msg.ReplyMarkup = b.GetButtons(message.Chat.ID) TODO:
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages["tj"]["Welcome"] + "\n\n" + messages.Messages["ru"]["Welcome"],
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(err)
	}
}
