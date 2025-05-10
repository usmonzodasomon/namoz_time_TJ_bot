package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
)

func (h *Handler) TaqvimHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	date := time.Now().Format("02.01.2006")
	taqvimTime, err := h.storage.GetTaqvimTime()
	if err != nil {
		log.Println(err)
		return
	}
	taqvimTime = h.GetTaqvimTimeForCurrentRegion(taqvimTime, user.RegionID)
	fmt.Println(taqvimTime)

	namazString := fmt.Sprintf(`
📆 <b><i>%s: %s, %s</i></b>
🏢 <b><i>%s:        %s</i></b>

<b><i>%s %s:</i></b>          <code>%s</code>
<b><i>%s %s:</i></b>              <code>%s</code>
<b><i>%s %s:</i></b>                  <code>%s</code>
<b><i>%s %s:</i></b>              <code>%s</code>
<b><i>%s %s:</i></b>          <code>%s</code>
`,
		messages.Messages[user.Language]["Today"], date, messages.Messages[user.Language][time.Now().Weekday().String()],
		messages.Messages[user.Language]["Region"], types.Regions[user.Language][user.RegionID-1],
		types.Stickers[0], types.NamazIndex[user.Language][0], taqvimTime.Fajr,
		types.Stickers[1], types.NamazIndex[user.Language][1], taqvimTime.Zuhr,
		types.Stickers[2], types.NamazIndex[user.Language][2], taqvimTime.Asr,
		types.Stickers[3], types.NamazIndex[user.Language][3], taqvimTime.Maghrib,
		types.Stickers[4], types.NamazIndex[user.Language][4], taqvimTime.Isha,
	)
	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        namazString,
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(err)
		return
	}
}
