package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/parser"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/storage"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/telegram"
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
	go s.UpdateTaqvimTime()
	go s.UpdateTime()
	go s.SendReminders()

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
		gocron.DurationJob(time.Minute),
		gocron.NewTask(
			s.SendReminders,
		),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)
	if err != nil {
		log.Println(err)
	}

	_, err = s.sh.NewJob(
		gocron.CronJob("5 0 * * *", false),
		gocron.NewTask(
			s.UpdateTaqvimTime,
		),
	)
	if err != nil {
		log.Println(err)
	}

	s.sh.Start()
}
