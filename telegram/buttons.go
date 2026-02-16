package telegram

import (
	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
)

func isLang(update *models.Update) bool {
	return update.Message != nil && (update.Message.Text == "🇹🇯 "+messages.Messages["ru"]["ChooseLanguageBtn"] ||
		update.Message.Text == "🇹🇯 "+messages.Messages["tj"]["ChooseLanguageBtn"])
}
func isTime(update *models.Update) bool {
	return update.Message != nil && (update.Message.Text == "🕓 "+messages.Messages["tj"]["NamazTimeBtn"] || update.Message.Text == "🕓 "+messages.Messages["ru"]["NamazTimeBtn"])
}

func isRegionButton(update *models.Update) bool {
	return update.Message != nil && (update.Message.Text == "🏙 "+messages.Messages["ru"]["ChooseRegionBtn"] || update.Message.Text == "🏙 "+messages.Messages["tj"]["ChooseRegionBtn"])
}

func isTaqvimButton(update *models.Update) bool {
	return update.Message != nil && (update.Message.Text == "🕌 "+messages.Messages["ru"]["Taqvim"] || update.Message.Text == "🕌 "+messages.Messages["tj"]["Taqvim"])
}

func isSettingsButton(update *models.Update) bool {
	return update.Message != nil && (update.Message.Text == "⚙️ "+messages.Messages["ru"]["SettingsBtn"] || update.Message.Text == "⚙️ "+messages.Messages["tj"]["SettingsBtn"])
}

func isRamadanButton(update *models.Update) bool {
	return update.Message != nil && (update.Message.Text == "🌙 "+messages.Messages["ru"]["RamadanBtn"] || update.Message.Text == "🌙 "+messages.Messages["tj"]["RamadanBtn"])
}
