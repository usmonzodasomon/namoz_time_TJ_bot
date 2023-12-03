package telegram

import (
	"echobot/messages"
	"echobot/types"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.CallbackQuery != nil {
			// Получаем данные о нажатой кнопке
			callbackData := update.CallbackQuery.Data
			userID := update.CallbackQuery.From.ID
			// messageID := update.CallbackQuery.Message.MessageID

			// Создаем новое сообщение как ответ на нажатие кнопки
			msg := tgbotapi.NewMessage(int64(userID), b.getMessage(int64(userID), "YourChoose")+": "+callbackData)
			if err := b.db.UpdateRegionID(update.CallbackQuery.Message.Chat.ID, types.RegionsID[callbackData]); err != nil {
				log.Println(err.Error())
			}
			// msg.ReplyToMessageID = messageID
			b.bot.Send(msg)

			// Отвечаем на CallbackQuery
			callbackResponse := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			b.bot.AnswerCallbackQuery(callbackResponse)
			continue
		}
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.processButton(update.Message)
		// b.handleMassage(update.Message)
	}
}

func (b *Bot) processButton(message *tgbotapi.Message) {
	switch message.Text {
	case "🕓 " + messages.Messages["ru"]["NamazTimeBtn"], "🕓 " + messages.Messages["tj"]["NamazTimeBtn"]:
		if err := b.time(message.Chat.ID); err != nil {
			log.Println(err.Error())
		}
	case "🇹🇯 " + messages.Messages["ru"]["ChooseLanguageBtn"], "🇹🇯 " + messages.Messages["tj"]["ChooseLanguageBtn"]:
		if err := b.language(message); err != nil {
			log.Println(err.Error())
		}
	case "🏙 " + messages.Messages["ru"]["ChooseRegionBtn"], "🏙 " + messages.Messages["tj"]["ChooseRegionBtn"]:
		if err := b.region(message); err != nil {
			log.Println(err.Error())
		}
	case "🇹🇯 Тоҷикӣ", "🇷🇺 Русский":
		if err := b.changeLanguage(message); err != nil {
			log.Print(err.Error())
		}
	default:
		b.Uncknown(message)
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		if err := b.start(message); err != nil {
			log.Println(err.Error())
		}
	case "time":
		if err := b.time(message.Chat.ID); err != nil {
			log.Println(err.Error())
		}
	case "region":
		if err := b.region(message); err != nil {
			log.Println(err.Error())
		}
	case "language":
		if err := b.language(message); err != nil {
			log.Println(err.Error())
		}
	default:
		b.Uncknown(message)
	}
}
