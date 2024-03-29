package scheduler

import (
	"echobot/parser"
	"echobot/storage"
	"echobot/telegram"
	"github.com/go-co-op/gocron/v2"
	"log"
	"time"
)

type Scheduler struct {
	sh       gocron.Scheduler
	parser   *parser.Parser
	storage  storage.Storage
	telegram *telegram.Bot
}

func NewScheduler(parser *parser.Parser, scheduler gocron.Scheduler, storage storage.Storage, telegram *telegram.Bot) *Scheduler {
	return &Scheduler{
		sh:       scheduler,
		parser:   parser,
		storage:  storage,
		telegram: telegram,
	}
}

func (s *Scheduler) Start() {
	s.UpdateTime()
	s.SendReminders()

	_, err := s.sh.NewJob(
		gocron.CronJob("0 */6 * * *", false),
		gocron.NewTask(
			s.UpdateTime,
		),
	)
	if err != nil {
		log.Println(err)
	}
	_, err = s.sh.NewJob(
		gocron.DurationJob(3*time.Minute),
		gocron.NewTask(
			s.SendReminders,
		),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)

	if err != nil {
		log.Println(err)
	}
	s.sh.Start()
}
