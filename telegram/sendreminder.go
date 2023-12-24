package telegram

import (
	"echobot/types"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var SendReminderTimeSleep = 5 * time.Minute

// Вызываем SendReminder каждые 5 минут чтобы найти пользоваталей в регионах которых настало время намаза
// Напоминание отправляются не за долго до наступления времени намаза (до 10 минут)

func (b *Bot) SendReminder() {
	for {
		for namazID := 0; namazID < 5; namazID++ {
			for regionID := 0; regionID < 13; regionID++ {
				if b.Parser.NamazTime == nil {
					break
				}
				time1 := getMinutes(b.Parser.NamazTime.Namaz[namazID].From) + types.RegionsTime[regionID]
				time2 := getMinutes(time.Now())
				// log.Println(time2, time1)
				if time1-time2 >= 0 && time1-time2 <= 10 && !types.SentNotifications[regionID][namazID] {
					if err := b.SendMessageForAllUsers(namazID, regionID); err != nil {
						log.Println(err.Error())
					}
					types.SentNotifications[regionID] = make(map[int]bool)
					types.SentNotifications[regionID][namazID] = true
					break
				}
			}
		}
		time.Sleep(SendReminderTimeSleep)
	}
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
			return err
		}

		msgID, err := b.db.GetLastMessageID(chatID)
		if err != nil {
			return err
		}

		if err := b.DeleteMessage(chatID, msgID); err != nil {
			log.Println(err.Error())
		}

		if namazID == 0 {
			if err := b.time(chatID); err != nil {
				return err
			}
		}

		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("*%s:* %s\n*%s:* %s %s %s %s",
			b.getMessage(chatID, "NextNamaz"),
			types.NamazIndex[lang][namazID],
			b.getMessage(chatID, "Time"),
			b.getMessage(chatID, "IntervalFrom"),
			b.Parser.NamazTime.Namaz[namazID].From.Format("15:04"),
			b.getMessage(chatID, "IntervalTo"),
			b.Parser.NamazTime.Namaz[namazID].To.Format("15:04")))
		msg.ParseMode = "Markdown"

		r, err := b.bot.Send(msg)
		if err != nil {
			log.Println("error sending message from timer func: ", err.Error())
		}
		if r.Chat == nil {
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

func (b *Bot) DeleteMessage(chatID int64, messageID int) error {
	if messageID == 0 {
		return nil
	}
	_, err := b.bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	})
	return err
}
