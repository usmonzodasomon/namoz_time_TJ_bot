package handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"log"
)

func (h *Handler) SettingsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	kb := inline.New(b).Row().
		Button("🇹🇯 "+messages.Messages[user.Language]["Language"], []byte("settings_language"), h.onInlineKeyboardSettingsLanguage).
		Button("🏙 "+messages.Messages[user.Language]["Region"], []byte("settings_region"), h.onInlineKeyboardSettingsRegion)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["Settings"] + ":",
		ReplyMarkup: &kb,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) onInlineKeyboardSettingsLanguage(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	if mes.Type == 1 {
		log.Println("MessageType is InaccessibleMessage")
		return
	}

	user, err := h.storage.GetUser(mes.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	kb := inline.New(b).Row().
		Button("🇹🇯 Тоҷикӣ", []byte("lang_tj"), h.onInlineKeyboardSelectLanguage).
		Button("🇷🇺 Русский", []byte("lang_ru"), h.onInlineKeyboardSelectLanguage)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      mes.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["ChooseLanguage"] + ":",
		ReplyMarkup: &kb,
	})
}

func (h *Handler) onInlineKeyboardSettingsRegion(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	if mes.Type == 1 {
		log.Println("MessageType is InaccessibleMessage")
		return
	}

	user, err := h.storage.GetUser(mes.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	kb := inline.New(b).Row()
	for i, region := range types.Regions[user.Language] {
		kb = kb.Button(region, []byte(region), h.onInlineKeyboardSelectRegion)
		if i%2 == 1 {
			kb = kb.Row()
		}
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      mes.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["ChooseRegion"] + ":",
		ReplyMarkup: &kb,
	})
}
