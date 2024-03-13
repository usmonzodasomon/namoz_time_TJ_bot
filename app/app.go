package app

import (
	"context"
	"echobot/scheduler"
	"echobot/telegram"
)

type App struct {
	bot       *telegram.Bot
	scheduler *scheduler.Scheduler
}

func NewApp(bot *telegram.Bot, scheduler *scheduler.Scheduler) *App {
	return &App{
		bot:       bot,
		scheduler: scheduler,
	}
}

func (a *App) Start(ctx context.Context) {
	go a.bot.Start(ctx)
	go a.scheduler.Start()
}
