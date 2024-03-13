package scheduler

import (
	"log"
	"time"
)

func (s *Scheduler) UpdateTime() {
	now := time.Now().Format("02.01.2006")

	namazTimeSl, err := s.parser.Parse(now[3:5], now[6:])
	if err != nil {
		log.Println(err)
		return
	}

	if len(namazTimeSl) == 0 {
		log.Println("empty namazTimeSl")
		return
	}
	if err := s.storage.UpdateNamazTime(namazTimeSl); err != nil {
		log.Println(err)
		return
	}

}
