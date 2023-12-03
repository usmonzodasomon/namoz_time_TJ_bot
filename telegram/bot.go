package telegram

import (
	"echobot/parser"
	"echobot/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	Parser *parser.Parser
	db     repository.Repository
}

func NewBot(bot *tgbotapi.BotAPI, parser *parser.Parser, db *repository.Sqlite) *Bot {
	return &Bot{
		bot:    bot,
		Parser: parser,
		db:     db,
	}
}

func (b *Bot) Start() error {
	go b.UpdateTime()
	go b.SendReminder()

	updates, err := b.GetUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)
	return nil
}

func (b *Bot) GetUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
