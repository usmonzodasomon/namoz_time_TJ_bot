package handler

import (
	"context"
	"echobot/messages"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strconv"
)

func (h *Handler) TaqvimHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		user.Language = "tj"
	}

	_, err = b.CopyMessage(ctx, &bot.CopyMessageParams{
		ChatID:      update.Message.Chat.ID,
		FromChatID:  strconv.FormatInt(1073064760, 10),
		MessageID:   269545,
		Caption:     messages.Messages[user.Language]["Taqvim"],
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(fmt.Sprintf("taqvimhandler error: %v", err))
		return
	}
}
