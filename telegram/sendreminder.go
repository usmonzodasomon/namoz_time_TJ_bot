package telegram

import (
	"echobot/types"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var SendReminderTimeSleep = 3 * time.Minute

// Вызываем SendReminder каждые 3 минут чтобы найти пользоваталей в регионах которых настало время намаза
// Напоминание отправляются не за долго до наступления времени намаза (до 10 минут)

func (b *Bot) SendRemindersProcedure() {
	for {
		b.SendReminders()
		time.Sleep(SendReminderTimeSleep)
	}
}

func (b *Bot) SendReminders() {
	if b.Parser.NamazTime == nil {
		return
	}

	for namazID := 0; namazID < 5; namazID++ {
		b.SendRemindersForNamaz(namazID)
	}
}

func (b *Bot) SendRemindersForNamaz(namazID int) {
	for regionID := 0; regionID < 13; regionID++ {
		b.SendRemindersForRegion(namazID, regionID)
	}
}

func (b *Bot) SendRemindersForRegion(namazID, regionID int) {
	nmzTimeForCurrRegion := getNamazTimeForCurrentRegionInMinutes(
		b.Parser.NamazTime.Namaz[namazID].From, regionID)

	nowInMinutes := getMinutes(time.Now())

	if isNamazTime(nmzTimeForCurrRegion, nowInMinutes) && !types.SendNotifications[regionID][namazID] {
		if err := b.SendMessageForAllUsers(namazID, regionID); err != nil {
			log.Println("error sending notification:", err)
		}
		types.SendNotifications[regionID] = make(map[int]bool)
		types.SendNotifications[regionID][namazID] = true
		return
	}
}

func isNamazTime(namazTimeForCurrentRegionInMinutes, nowInMinutes int) bool {
	return namazTimeForCurrentRegionInMinutes-nowInMinutes >= 0 &&
		namazTimeForCurrentRegionInMinutes-nowInMinutes <= 10
}

func getNamazTimeForCurrentRegionInMinutes(namazTime time.Time, regionID int) int {
	return getMinutes(namazTime) + types.RegionsTime[regionID]
}

func getMinutes(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

func (b *Bot) SendMessageForAllUsers(namazID, regionID int) error {
	usersChatIDs, err := b.db.GetAllUsersByRegionID(regionID)
	if err != nil {
		return err
	}

	for _, chatID := range usersChatIDs {
		lang, err := b.db.GetLang(chatID)
		if err != nil {
			log.Println("Error getting language: ", err, chatID)
			continue
		}

		msgID, err := b.db.GetLastMessageID(chatID)
		if err != nil {
			log.Println("Error getting last message ID: ", err, chatID)
			continue
		}

		if err := b.DeleteMessage(chatID, msgID); err != nil {
			log.Println("Error deleting message: ", err)
		}

		if namazID == 0 {
			if err := b.time(chatID); err != nil {
				log.Println("Error sending common message: ", err, chatID)
				continue
			}
		}

		msg := b.getNextNamazMessage(chatID, lang, namazID, regionID)
		r, err := b.bot.Send(msg)
		if err != nil {
			log.Println("error sending next namaz time message : ", err.Error())
		}

		if r.Chat == nil { // user bloks bot
			if err := b.db.DeleteUser(chatID); err != nil {
				log.Println(err.Error())
			}
			continue
		}

		if err := b.db.UpdateLastMessageID(r.Chat.ID, r.MessageID); err != nil {
			log.Println("error updating message id: ", err.Error())
		}
	}
	return nil
}

func (b *Bot) getNextNamazMessage(chatID int64, lang string, namazID int, region_ID int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("*%s:* %s\n*%s:* %s %s %s %s",
		b.getMessage(chatID, "NextNamaz"),
		types.NamazIndex[lang][namazID],
		b.getMessage(chatID, "Time"),
		b.getMessage(chatID, "IntervalFrom"),
		b.Parser.NamazTime.Namaz[namazID].From.
			Add(time.Minute*time.Duration(types.RegionsTime[region_ID])).
			Format("15:04"), // находим время намаза для региона
		b.getMessage(chatID, "IntervalTo"),
		b.Parser.NamazTime.Namaz[namazID].To.
			Add(time.Minute*time.Duration(types.RegionsTime[region_ID])).
			Format("15:04"))) // находим время истечения намаза для региона
	msg.ParseMode = "Markdown"
	return msg
}

func (b *Bot) DeleteMessage(chatID int64, messageID int) error {
	if messageID == 0 {
		return nil
	}
	_, err := b.bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	})
	if err != nil {
		return fmt.Errorf("error deleting message: %v, chat_id: %v, message_id = %v", err, chatID, messageID)
	}
	return nil
}
