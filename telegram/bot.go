package telegram

import (
	"context"

	"github.com/go-telegram/bot"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/handler"
)

type Bot struct {
	Bot     *bot.Bot
	Handler *handler.Handler
}

func NewBot(bot *bot.Bot, handler *handler.Handler) *Bot {
	return &Bot{
		Bot:     bot,
		Handler: handler,
	}
}

func (b *Bot) Start(ctx context.Context) {
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, b.Handler.StartHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/time", bot.MatchTypeExact, b.Handler.TimeHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/language", bot.MatchTypeExact, b.Handler.LangHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/region", bot.MatchTypeExact, b.Handler.RegionHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/settings", bot.MatchTypeExact, b.Handler.SettingsHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/taqvim", bot.MatchTypeExact, b.Handler.TaqvimHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/mailing_me", bot.MatchTypeExact, b.Handler.MailingMeHandler, b.Handler.AuthHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/mailing_all", bot.MatchTypeExact, b.Handler.MailingAllHandler, b.Handler.AuthHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/stat", bot.MatchTypeExact, b.Handler.StatHandler, b.Handler.AuthHandler)
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/backup", bot.MatchTypeExact, b.Handler.BackupHandler, b.Handler.AuthHandler)

	b.Bot.RegisterHandlerMatchFunc(isLang, b.Handler.LangHandler)
	b.Bot.RegisterHandlerMatchFunc(isTime, b.Handler.TimeHandler)
	b.Bot.RegisterHandlerMatchFunc(isLangButton, b.Handler.ChangeLanguage)
	b.Bot.RegisterHandlerMatchFunc(isRegionButton, b.Handler.RegionHandler)
	b.Bot.RegisterHandlerMatchFunc(isSettingsButton, b.Handler.SettingsHandler)
	b.Bot.RegisterHandlerMatchFunc(isTaqvimButton, b.Handler.TaqvimHandler)

	b.Bot.Start(ctx)
}
