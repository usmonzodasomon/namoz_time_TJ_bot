package handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"log"
	"time"
)

func (h *Handler) TimeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	date := time.Now().Format("02.01.2006")
	namazTime, err := h.storage.GetNamazTime(date)
	if err != nil {
		log.Println(err)
		return
	}
	namazTime = h.GetNamazTimeForCurrentRegion(namazTime, user.RegionID)
	fmt.Println(namazTime)

	namazString := fmt.Sprintf(`
📆 <b><i>%s: %s, %s</i></b>
🏢 <b><i>%s:        %s</i></b>
📰 <b><i>Манбаъ: shuroiulamo.tj</i></b>

<b><i>%s %s:</i></b>          <code>%s - %s</code>
<b><i>%s %s:</i></b>              <code>%s - %s</code>
<b><i>%s %s:</i></b>                  <code>%s - %s</code>
<b><i>%s %s:</i></b>              <code>%s - %s</code>
<b><i>%s %s:</i></b>          <code>%s - %s</code>
`,
		messages.Messages[user.Language]["Today"], date, messages.Messages[user.Language][time.Now().Weekday().String()],
		messages.Messages[user.Language]["Region"], types.Regions[user.Language][user.RegionID-1],
		types.Stickers[0], types.NamazIndex[user.Language][0], namazTime.FajrFrom, namazTime.FajrTo,
		types.Stickers[1], types.NamazIndex[user.Language][1], namazTime.ZuhrFrom, namazTime.ZuhrTo,
		types.Stickers[2], types.NamazIndex[user.Language][2], namazTime.AsrFrom, namazTime.AsrTo,
		types.Stickers[3], types.NamazIndex[user.Language][3], namazTime.MaghribFrom, namazTime.MaghribTo,
		types.Stickers[4], types.NamazIndex[user.Language][4], namazTime.IshaFrom, namazTime.IshaTo,
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
