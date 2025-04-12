package scheduler

import (
	"github.com/usmonzodasomon/namoz_time_TJ_bot/parser"
	"log"
)

func (s *Scheduler) UpdateTaqvimTime() {
	times, err := parser.GetTaqvimNamazTime()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	
	if err := s.storage.UpdateTaqvimTime(*times); err != nil {
		log.Println(err)
		return
	}
}
