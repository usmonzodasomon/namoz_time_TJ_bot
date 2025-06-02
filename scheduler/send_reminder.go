package scheduler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
)

var (
	ErrBlockUser      = errors.New(`forbidden, Forbidden: bot was blocked by the user`)
	ErrDeactivateUser = errors.New(`forbidden, Forbidden: user is deactivated`)
	ErrChatNotFound   = errors.New(`bad request, Bad Request: chat not found`)
)

func (s *Scheduler) SendReminders() {
	taqvimTime, err := s.storage.GetTaqvimTime()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for namazID := 0; namazID < 5; namazID++ {
		s.SendRemindersForNamaz(namazID, taqvimTime)
	}
}

func (s *Scheduler) SendRemindersForNamaz(namazID int, taqvimTime types.TaqvimTime) {
	for regionID := 1; regionID <= 14; regionID++ {
		s.SendRemindersForRegion(namazID, regionID, taqvimTime)
	}
}

func (s *Scheduler) SendRemindersForRegion(namazID, regionID int, taqvimTime types.TaqvimTime) {
	nowInMinutes := getMinutes(time.Now())

	taqvimTimeStr := getTaqvimTimeWithNamazID(taqvimTime, namazID)
	taqvimTimeObj, _ := time.Parse("15:04", taqvimTimeStr)
	nmzTimeForCurrRegion := getMinutes(taqvimTimeObj.Add(time.Minute * time.Duration(types.RegionsTime[regionID])))

	if isNamazTime(nmzTimeForCurrRegion, nowInMinutes) && !types.SendNotifications[regionID][namazID] {
		if err := s.SendMessageForAllUsers(namazID, regionID, taqvimTime); err != nil {
			log.Println("error sending notification:", err)
		}
		types.SendNotifications[regionID] = make(map[int]bool)
		types.SendNotifications[regionID][namazID] = true
	}
}

func getTaqvimTimeWithNamazID(taqvimTime types.TaqvimTime, namazID int) string {
	switch namazID {
	case 0:
		return taqvimTime.Fajr
	case 1:
		return taqvimTime.Zuhr
	case 2:
		return taqvimTime.Asr
	case 3:
		return taqvimTime.Maghrib
	case 4:
		return taqvimTime.Isha
	}
	return ""
}

func isNamazTime(namazTimeForCurrentRegionInMinutes, nowInMinutes int) bool {
	return namazTimeForCurrentRegionInMinutes-nowInMinutes >= 0 &&
		namazTimeForCurrentRegionInMinutes-nowInMinutes <= 10
}

func getMinutes(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

func (s *Scheduler) SendMessageForAllUsers(namazID, regionID int, taqvimTime types.TaqvimTime) error {
	users, err := s.storage.GetAllUsersByRegionID(regionID)
	if err != nil {
		return err
	}

	ch := make(chan types.User)
	wg := &sync.WaitGroup{}
	go func(ch chan types.User) {
		for _, user := range users {
			ch <- user
		}
		close(ch)
	}(ch)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for user := range ch {
				s.SendMessageForUser(user, namazID, regionID, taqvimTime)
			}
		}()
	}
	wg.Wait()
	return nil
}

func (s *Scheduler) SendMessageForUser(user types.User, namazID, regionID int, taqvimTime types.TaqvimTime) {
	if err := s.DeleteMessage(user.ChatID, user.LastMessageID); err != nil {
		log.Println("Error deleting message: ", err)
	}

	if namazID == 0 {
		s.telegram.Handler.TimeHandler(context.Background(), s.telegram.Bot, &models.Update{
			Message: &models.Message{
				Chat: models.Chat{
					ID: user.ChatID,
				},
			},
		})
	}

	r, err := s.telegram.Bot.SendMessage(context.Background(), s.getNextNamazMessage(user, namazID, regionID, taqvimTime))
	if err != nil {
		log.Println("error sending next namaz time message : ", err.Error())
		if err.Error() == ErrBlockUser.Error() ||
			err.Error() == ErrDeactivateUser.Error() ||
			err.Error() == ErrChatNotFound.Error() {
			log.Println("deleting user: ", user.ChatID)
			if err := s.storage.DeleteUser(user.ChatID); err != nil {
				log.Println(err.Error())
			}
		}
		return
	}

	if err := s.storage.UpdateUser(types.User{
		ChatID:        user.ChatID,
		LastMessageID: r.ID,
	}); err != nil {
		log.Println("error updating message id: ", err.Error())
	}
}

func (s *Scheduler) getNextNamazMessage(user types.User, namazID, regionID int, taqvimTime types.TaqvimTime) *bot.SendMessageParams {
	timeStr := getTaqvimTimeWithNamazID(taqvimTime, namazID)
	adjustedTime, _ := time.Parse("15:04", timeStr)
	adjustedTime = adjustedTime.Add(time.Minute * time.Duration(types.RegionsTime[regionID]))

	msg := fmt.Sprintf(
		`<b>%s:</b> %s
<b>%s:</b> %s`,
		messages.Messages[user.Language]["NextNamaz"],
		types.NamazIndex[user.Language][namazID],
		messages.Messages[user.Language]["Time"],
		adjustedTime.Format("15:04"))

	return &bot.SendMessageParams{
		ChatID:    user.ChatID,
		Text:      msg,
		ParseMode: models.ParseModeHTML,
	}
}

func (s *Scheduler) DeleteMessage(chatID int64, messageID int) error {
	if messageID == 0 {
		return nil
	}
	_, err := s.telegram.Bot.DeleteMessage(context.Background(), &bot.DeleteMessageParams{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		return fmt.Errorf("error deleting message: %v, chat_id: %v, message_id = %v", err, chatID, messageID)
	}
	return nil
}
