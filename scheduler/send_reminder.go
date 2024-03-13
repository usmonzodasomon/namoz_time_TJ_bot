package scheduler

import (
	"context"
	"echobot/messages"
	"echobot/types"
	"errors"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"time"
)

var ErrBlockUser = errors.New(`unexpected response statusCode 403 for method sendMessage, {"ok":false,"error_code":403,"description":"Forbidden: bot was blocked by the user"}`)

func (s *Scheduler) SendReminders() {
	date := time.Now().Format("02.01.2006")
	namazTime, err := s.storage.GetNamazTime(date)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for namazID := 0; namazID < 5; namazID++ {
		s.SendRemindersForNamaz(namazID, namazTime)
	}
}

func (s *Scheduler) SendRemindersForNamaz(namazID int, namazTime types.NamazTime) {
	for regionID := 1; regionID <= 14; regionID++ {
		s.SendRemindersForRegion(namazID, regionID, namazTime)
	}

}

func (s *Scheduler) SendRemindersForRegion(namazID, regionID int, namazTime types.NamazTime) {
	namazTime = s.telegram.Handler.GetNamazTimeForCurrentRegion(namazTime, regionID)

	nowInMinutes := getMinutes(time.Now())
	nmzTimeForCurrRegion := getMinutes(getNamazTimeWithNamazID(namazTime, namazID).From.
		Add(time.Minute * time.Duration(types.RegionsTime[regionID])))

	if isNamazTime(nmzTimeForCurrRegion, nowInMinutes) && !types.SendNotifications[regionID][namazID] {
		if err := s.SendMessageForAllUsers(namazID, regionID, namazTime); err != nil {
			log.Println("error sending notification:", err)
		}
		types.SendNotifications[regionID] = make(map[int]bool)
		types.SendNotifications[regionID][namazID] = true
	}
}

func getNamazTimeWithNamazID(namazTime types.NamazTime, namazID int) types.NamazTimeStruct {
	switch namazID {
	case 0:
		FajrFrom, err := time.Parse("15:04", namazTime.FajrFrom)
		if err != nil {
			log.Println(err)
		}
		FajrTo, err := time.Parse("15:04", namazTime.FajrTo)
		if err != nil {
			log.Println(err)
		}
		return types.NamazTimeStruct{
			From: FajrFrom,
			To:   FajrTo,
		}
	case 1:
		ZuhrFrom, err := time.Parse("15:04", namazTime.ZuhrFrom)
		if err != nil {
			log.Println(err)
		}
		ZuhrTo, err := time.Parse("15:04", namazTime.ZuhrTo)
		if err != nil {
			log.Println(err)
		}
		return types.NamazTimeStruct{
			From: ZuhrFrom,
			To:   ZuhrTo,
		}
	case 2:
		AsrFrom, err := time.Parse("15:04", namazTime.AsrFrom)
		if err != nil {
			log.Println(err)
		}
		AsrTo, err := time.Parse("15:04", namazTime.AsrTo)
		if err != nil {
			log.Println(err)
		}
		return types.NamazTimeStruct{
			From: AsrFrom,
			To:   AsrTo,
		}
	case 3:
		MaghribFrom, err := time.Parse("15:04", namazTime.MaghribFrom)
		if err != nil {
			log.Println(err)
		}
		MaghribTo, err := time.Parse("15:04", namazTime.MaghribTo)
		if err != nil {
			log.Println(err)
		}
		return types.NamazTimeStruct{
			From: MaghribFrom,
			To:   MaghribTo,
		}
	case 4:
		IshaFrom, err := time.Parse("15:04", namazTime.IshaFrom)
		if err != nil {
			log.Println(err)
		}
		IshaTo, err := time.Parse("15:04", namazTime.IshaTo)
		if err != nil {
			log.Println(err)
		}
		return types.NamazTimeStruct{
			From: IshaFrom,
			To:   IshaTo,
		}
	}
	return types.NamazTimeStruct{}
}

func isNamazTime(namazTimeForCurrentRegionInMinutes, nowInMinutes int) bool {
	return namazTimeForCurrentRegionInMinutes-nowInMinutes >= 0 &&
		namazTimeForCurrentRegionInMinutes-nowInMinutes <= 10
}

func getMinutes(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

func (s *Scheduler) SendMessageForAllUsers(namazID, regionID int, namazTime types.NamazTime) error {
	users, err := s.storage.GetAllUsersByRegionID(regionID)
	if err != nil {
		return err
	}
	fmt.Println(users)

	for _, user := range users {
		s.SendMessageForUser(user, namazID, regionID, namazTime)
	}
	return nil
}

func (s *Scheduler) SendMessageForUser(user types.User, namazID, regionID int, namazTime types.NamazTime) {
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

	r, err := s.telegram.Bot.SendMessage(context.Background(), s.getNextNamazMessage(user, namazID, regionID, namazTime))
	if err != nil {
		log.Println("error sending next namaz time message : ", err.Error())
		if err.Error() == ErrBlockUser.Error() {
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

func (s *Scheduler) getNextNamazMessage(user types.User, namazID, regionID int, namazTime types.NamazTime) *bot.SendMessageParams {
	msg := fmt.Sprintf(
		`<b>%s:</b> %s
<b>%s:</b> %s %s %s %s`,
		messages.Messages[user.Language]["NextNamaz"],
		types.NamazIndex[user.Language][namazID],
		messages.Messages[user.Language]["Time"],
		messages.Messages[user.Language]["IntervalFrom"],
		getNamazTimeWithNamazID(namazTime, namazID).From.
			Add(time.Minute*time.Duration(types.RegionsTime[regionID])).
			Format("15:04"), // находим время намаза для региона
		messages.Messages[user.Language]["IntervalTo"],
		getNamazTimeWithNamazID(namazTime, namazID).To.
			Add(time.Minute*time.Duration(types.RegionsTime[regionID])).
			Format("15:04")) // находим время истечения намаза для региона
	if namazID == 0 {
		msg = fmt.Sprintf("<b>🌙  🌙  🌙  🌙 🌙 🌙  🌙  🌙  🌙 🌙 🌙\n\n%s: %s</b>\n\n🌙  🌙  🌙  🌙 🌙 🌙  🌙  🌙  🌙 🌙 🌙\n\n", messages.Messages[user.Language]["Suhur"], "Валисавми ғаддин мин шаҳри  рамазоналлазӣ фаризатан навайту.") + msg
	} else if namazID == 3 {
		msg = fmt.Sprintf("<b>🌙  🌙  🌙  🌙 🌙 🌙  🌙  🌙  🌙 🌙 🌙\n\n%s: %s</b>\n\n🌙  🌙  🌙  🌙 🌙 🌙  🌙  🌙  🌙 🌙 🌙\n\n", messages.Messages[user.Language]["Iftar"], "Аллоҳума лака сумту ва бика оманту ва алайка таваккалту ва ало ризқика афтарту.") + msg

	}
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
