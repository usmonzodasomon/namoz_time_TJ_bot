package telegram

import (
	"echobot/messages"
	"github.com/go-telegram/bot/models"
)

func isLang(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	return update.Message.Text == "🇹🇯 "+messages.Messages["ru"]["ChooseLanguageBtn"] ||
		update.Message.Text == "🇹🇯 "+messages.Messages["tj"]["ChooseLanguageBtn"]
}
func isTime(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	return update.Message.Text == "🕓 "+messages.Messages["tj"]["NamazTimeBtn"] || update.Message.Text == "🕓 "+messages.Messages["ru"]["NamazTimeBtn"]
}

func isLangButton(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	return update.Message.Text == "🇹🇯 Тоҷикӣ" || update.Message.Text == "🇷🇺 Русский"
}

func isRegionButton(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	return update.Message.Text == "🏙 "+messages.Messages["ru"]["ChooseRegionBtn"] || update.Message.Text == "🏙 "+messages.Messages["tj"]["ChooseRegionBtn"]
}

func isTaqvimButton(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	return update.Message.Text == "🕌 "+messages.Messages["tj"]["Taqvim"] || update.Message.Text == "🕌 "+messages.Messages["ru"]["Taqvim"]
}
