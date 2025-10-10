package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/storage"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/telegram"
	"log"
	"time"
)

type Scheduler struct {
	sh          gocron.Scheduler
	storage     storage.Storage
	telegram    *telegram.Bot
	adminChatID int64
	dbHost      string
	dbUser      string
	dbName      string
	dbPassword  string
}

func NewScheduler(scheduler gocron.Scheduler, storage storage.Storage, telegram *telegram.Bot, adminChatID int64, dbHost, dbUser, dbName, dbPassword string) *Scheduler {
	return &Scheduler{
		sh:          scheduler,
		storage:     storage,
		telegram:    telegram,
		adminChatID: adminChatID,
		dbHost:      dbHost,
		dbUser:      dbUser,
		dbName:      dbName,
		dbPassword:  dbPassword,
	}
}

func (s *Scheduler) Start() {
	go s.UpdateTaqvimTime()
	go s.UpdateTime()
	go s.SendReminders()

	_, err := s.sh.NewJob(gocron.CronJob("0 */6 * * *", false), gocron.NewTask(s.UpdateTime))
	if err != nil {
		log.Println(err)
	}

	_, err = s.sh.NewJob(gocron.DurationJob(time.Minute), gocron.NewTask(s.SendReminders), gocron.WithSingletonMode(gocron.LimitModeReschedule))
	if err != nil {
		log.Println(err)
	}

	_, err = s.sh.NewJob(gocron.CronJob("5 0 * * *", false), gocron.NewTask(s.UpdateTaqvimTime))
	if err != nil {
		log.Println(err)
	}

	// Database backup job - runs daily at 03:00
	_, err = s.sh.NewJob(gocron.CronJob("0 3 * * *", false), gocron.NewTask(s.SendDatabaseBackup))
	if err != nil {
		log.Println(err)
	}

	s.sh.Start()
}
