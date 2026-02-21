package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
)

const (
	ramadanSuhur = 0
	ramadanIftar = 1
)

func (s *Scheduler) SendRamadanReminders() {
	today := time.Now().Format("2006-01-02")
	ramadanTime, err := s.storage.GetRamadanTimeByDate(today)
	if err != nil {
		return
	}

	for regionID := 1; regionID <= 19; regionID++ {
		s.sendRamadanReminderForRegion(ramadanSuhur, regionID, ramadanTime)
		s.sendRamadanReminderForRegion(ramadanIftar, regionID, ramadanTime)
	}
}

func (s *Scheduler) sendRamadanReminderForRegion(reminderType, regionID int, ramadanTime types.RamadanTime) {
	if types.SendRamadanNotifications[regionID] != nil && types.SendRamadanNotifications[regionID][reminderType] {
		return
	}

	var timeStr string
	if reminderType == ramadanSuhur {
		timeStr = ramadanTime.SubhSadiq
	} else {
		timeStr = ramadanTime.Shom
	}

	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		log.Println("error parsing ramadan time:", err)
		return
	}
	adjustedTime := parsedTime.Add(time.Minute * time.Duration(types.RegionsTime[regionID]))
	adjustedMinutes := getMinutes(adjustedTime)
	nowMinutes := getMinutes(time.Now())

	if !isNamazTime(adjustedMinutes, nowMinutes) {
		return
	}

	users, err := s.storage.GetAllUsersByRegionID(regionID)
	if err != nil {
		log.Println("error getting users for ramadan reminder:", err)
		return
	}

	ch := make(chan types.User)
	wg := &sync.WaitGroup{}
	go func() {
		for _, user := range users {
			ch <- user
		}
		close(ch)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range ch {
				s.sendRamadanMessageToUser(user, reminderType, adjustedTime)
			}
		}()
	}
	wg.Wait()

	types.SendRamadanNotifications[regionID] = make(map[int]bool)
	types.SendRamadanNotifications[regionID][reminderType] = true
}

func (s *Scheduler) sendRamadanMessageToUser(user types.User, reminderType int, adjustedTime time.Time) {
	if err := s.DeleteMessage(user.ChatID, user.LastRamadanMsgID); err != nil {
		log.Println("error deleting ramadan message:", err)
	}

	var msg string
	timeFormatted := adjustedTime.Format("15:04")

	if reminderType == ramadanSuhur {
		if user.Language == "tj" {
			msg = fmt.Sprintf("🌙 <b>Субҳи содиқ соати %s.</b>\n\n<i>%s: Валисавми ғаддин мин шаҳри рамазоналлазӣ фаризатан навайту.</i>",
				timeFormatted, messages.Messages[user.Language]["Suhur"])
		} else {
			msg = fmt.Sprintf("🌙 <b>Время сухура до %s.</b>\n\n<i>%s: Валисавми ғаддин мин шаҳри рамазоналлазӣ фаризатан навайту.</i>",
				timeFormatted, messages.Messages[user.Language]["Suhur"])
		}
	} else {
		if user.Language == "tj" {
			msg = fmt.Sprintf("🌙 <b>Вақти ифтор соати %s.</b>\n\n<i>%s: Аллоҳума лака сумту ва бика оманту ва алайка таваккалту ва ало ризқика афтарту.</i>",
				timeFormatted, messages.Messages[user.Language]["Iftar"])
		} else {
			msg = fmt.Sprintf("🌙 <b>Время ифтара в %s.</b>\n\n<i>%s: Аллоҳума лака сумту ва бика оманту ва алайка таваккалту ва ало ризқика афтарту.</i>",
				timeFormatted, messages.Messages[user.Language]["Iftar"])
		}
	}

	r, err := s.telegram.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID:    user.ChatID,
		Text:      msg,
		ParseMode: models.ParseModeHTML,
	})
	if err != nil {
		log.Println("error sending ramadan reminder:", err)
		if err.Error() == ErrBlockUser.Error() ||
			err.Error() == ErrDeactivateUser.Error() ||
			err.Error() == ErrChatNotFound.Error() {
			log.Println("deleting user:", user.ChatID)
			if err := s.storage.DeleteUser(user.ChatID); err != nil {
				log.Println(err.Error())
			}
		}
		return
	}

	if err := s.storage.UpdateUser(types.User{
		ChatID:           user.ChatID,
		LastRamadanMsgID: r.ID,
	}); err != nil {
		log.Println("error updating ramadan message id:", err)
	}
}
