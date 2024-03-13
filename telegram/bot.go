package telegram

import (
	"context"
	"echobot/handler"
	"github.com/go-telegram/bot"
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
	b.Bot.RegisterHandler(bot.HandlerTypeMessageText, "/taqvim", bot.MatchTypeExact, b.Handler.TaqvimHandler)

	b.Bot.RegisterHandlerMatchFunc(isLang, b.Handler.LangHandler)
	b.Bot.RegisterHandlerMatchFunc(isTime, b.Handler.TimeHandler)
	b.Bot.RegisterHandlerMatchFunc(isLangButton, b.Handler.ChangeLanguage)
	b.Bot.RegisterHandlerMatchFunc(isRegionButton, b.Handler.RegionHandler)
	b.Bot.RegisterHandlerMatchFunc(isTaqvimButton, b.Handler.TaqvimHandler)

	b.Bot.Start(ctx)
	//go b.UpdateTimeProcedure()
	//go b.SendRemindersProcedure()
	//
	//updates, err := b.GetUpdatesChannel()
	//if err != nil {
	//	return err
	//}
	//
	//b.handleUpdates(updates)
	//return nil
}
