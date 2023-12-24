package telegram

import (
	"echobot/messages"
	"echobot/types"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) start(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.Messages["tj"]["Welcome"]+"\n\n"+messages.Messages["ru"]["Welcome"])
	user := types.User{
		ChatID:   message.Chat.ID,
		RegionID: 0,
		Username: message.From.UserName,
		Language: "tj",
	}
	if err := b.db.CreateUser(user); err != nil {
		if err.Error() != "UNIQUE constraint failed: users.chat_id" {
			return err
		}
	}

	msg.ReplyMarkup = b.GetButtons(message.Chat.ID)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) time(chatID int64) error {
	regionID, err := b.db.GetRegionID(chatID)
	if err != nil {
		return err
	}

	minute := types.RegionsTime[int(regionID)]

	if b.Parser.NamazTime == nil {
		return errors.New("error parsing namaz time")
	}
	namazTime := getNamazTimeForCurrentUser(*b.Parser.NamazTime, minute)

	namazString := fmt.Sprintf("📆 _*%s: %s*_\n\n",
		b.getMessage(chatID, "Today"),
		strings.Replace(namazTime.Today.Format("02.01.2006"), ".", "\\.", -1))

	for i := 0; i < 5; i++ {
		lang, err := b.db.GetLang(chatID)
		if err != nil {
			return err
		}
		namazString += fmt.Sprintf("*_%s %s:_*          `%s %s %s %s`\n",
			types.Stickers[i],
			types.NamazIndex[lang][i],
			b.getMessage(chatID, "IntervalFrom"),
			namazTime.Namaz[i].From.Format("15:04"),
			b.getMessage(chatID, "IntervalTo"),
			namazTime.Namaz[i].To.Format("15:04"))
	}
	msg := tgbotapi.NewMessage(chatID, namazString)
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = b.GetButtons(chatID)
	_, err = b.bot.Send(msg)
	return err
}

func getNamazTimeForCurrentUser(v types.NamazTime, minute int) types.NamazTime {
	duration := time.Duration(minute) * time.Minute
	for i := range v.Namaz {
		v.Namaz[i].From = v.Namaz[i].From.Add(duration)
		v.Namaz[i].To = v.Namaz[i].To.Add(duration)
	}
	return v
}

func (b *Bot) region(message *tgbotapi.Message) error {
	buttons := make([][]tgbotapi.InlineKeyboardButton, 0)
	lang, err := b.db.GetLang(message.Chat.ID)
	if err != nil {
		return err
	}
	for i, region := range types.Regions[lang] {
		if i%2 == 0 {
			buttonsRow := tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(region, region),
			)
			if i+1 < len(types.Regions[lang]) {
				buttonsRow = append(buttonsRow, tgbotapi.NewInlineKeyboardButtonData(types.Regions[lang][i+1], types.Regions[lang][i+1]))
			}
			buttons = append(buttons, buttonsRow)
		}
	}

	replyMarkup := tgbotapi.NewInlineKeyboardMarkup(buttons...)

	msg := tgbotapi.NewMessage(message.Chat.ID, b.getMessage(message.Chat.ID, "ChooseRegion")+":")
	msg.ReplyMarkup = replyMarkup
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) Uncknown(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.getMessage(message.Chat.ID, "UnknownCommand"))
	msg.ReplyMarkup = b.GetButtons(message.Chat.ID)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) language(message *tgbotapi.Message) error {
	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🇹🇯 Тоҷикӣ"),
			tgbotapi.NewKeyboardButton("🇷🇺 Русский"),
		),
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, b.getMessage(message.Chat.ID, "ChooseLanguage")+": ")
	msg.ReplyMarkup = replyKeyboard
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) changeLanguage(message *tgbotapi.Message) error {
	lang := "ru"
	if message.Text == "🇹🇯 Тоҷикӣ" {
		lang = "tj"
	}
	if message.Text == "🇷🇺 Русский" {
		lang = "ru"
	}
	if err := b.db.UpdateLanguage(message.Chat.ID, lang); err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, messages.Messages[lang]["YourChoose"]+": "+message.Text)

	lang, err := b.db.GetLang(message.Chat.ID)
	if err != nil {
		return err
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
	msg.ReplyMarkup = replyKeyboard

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) getMessage(chatID int64, message string) string {
	lang, err := b.db.GetLang(chatID)
	if err != nil {
		log.Println(fmt.Errorf("error getting message, message is: %s, err: %s", message, err.Error()))
	}
	return messages.Messages[lang][message]
}
