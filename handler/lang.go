package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"log"
)

func (h *Handler) LangHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "🇹🇯 Тоҷикӣ"},
				{Text: "🇷🇺 Русский"},
			},
		},
		ResizeKeyboard: true,
		Selective:      true,
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["ChooseLanguage"] + ":",
		ReplyMarkup: kb,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) ChangeLanguage(ctx context.Context, b *bot.Bot, update *models.Update) {
	lang := "tj"
	if update.Message.Text == "🇹🇯 Тоҷикӣ" {
		lang = "tj"
	}
	if update.Message.Text == "🇷🇺 Русский" {
		lang = "ru"
	}

	if err := h.storage.UpdateUser(types.User{
		ChatID:   update.Message.Chat.ID,
		Language: lang,
	}); err != nil {
		log.Println(err)
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages[lang]["YourChoose"] + ": " + update.Message.Text,
		ReplyMarkup: inlineButtonMain(lang),
	})
	if err != nil {
		log.Println(err)
		return
	}
}
