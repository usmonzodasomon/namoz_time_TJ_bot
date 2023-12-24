package telegram

import (
	"echobot/messages"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) GetButtons(chatID int64) tgbotapi.ReplyKeyboardMarkup {
	lang := "tj"
	lang, err := b.db.GetLang(chatID)
	if err != nil {
		log.Println("error getting language: ", err)
	}
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🕓 "+messages.Messages[lang]["NamazTimeBtn"]),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🇹🇯 "+messages.Messages[lang]["ChooseLanguageBtn"]),
			tgbotapi.NewKeyboardButton("🏙 "+messages.Messages[lang]["ChooseRegionBtn"]),
		),
	)
	return replyKeyboard
}
