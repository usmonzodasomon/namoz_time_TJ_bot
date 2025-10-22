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

func (h *Handler) LangHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	kb := inline.New(b).Row().
		Button("🇹🇯 Тоҷикӣ", []byte("lang_tj"), h.onInlineKeyboardSelectLanguage).
		Button("🇷🇺 Русский", []byte("lang_ru"), h.onInlineKeyboardSelectLanguage)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["ChooseLanguage"] + ":",
		ReplyMarkup: &kb,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) onInlineKeyboardSelectLanguage(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	if mes.Type == 1 {
		log.Println("MessageType is InaccessibleMessage")
		return
	}

	lang := string(data)
	if lang == "lang_tj" {
		lang = "tj"
	} else if lang == "lang_ru" {
		lang = "ru"
	}

	if err := h.storage.UpdateUser(types.User{
		ChatID:   mes.Message.Chat.ID,
		Language: lang,
	}); err != nil {
		log.Println(err)
		return
	}

	langDisplay := "🇹🇯 Тоҷикӣ"
	if lang == "ru" {
		langDisplay = "🇷🇺 Русский"
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      mes.Message.Chat.ID,
		Text:        messages.Messages[lang]["YourChoose"] + ": " + langDisplay,
		ReplyMarkup: inlineButtonMain(lang),
	})
}
