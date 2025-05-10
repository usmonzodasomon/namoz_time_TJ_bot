package handler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode/utf8"

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
	weekday := messages.Messages[user.Language][time.Now().Weekday().String()]
	region := types.Regions[user.Language][user.RegionID-1]

	taqvimTime, err := h.storage.GetTaqvimTime()
	if err != nil {
		log.Println(err)
		return
	}
	taqvimTime = h.GetTaqvimTimeForCurrentRegion(taqvimTime, user.RegionID)

	type Entry struct {
		Emoji string
		Name  string
		Time  string
	}
	entries := []Entry{
		{types.Stickers[0], types.NamazIndex[user.Language][0], taqvimTime.Fajr},
		{types.Stickers[1], types.NamazIndex[user.Language][1], taqvimTime.Zuhr},
		{types.Stickers[2], types.NamazIndex[user.Language][2], taqvimTime.Asr},
		{types.Stickers[3], types.NamazIndex[user.Language][3], taqvimTime.Maghrib},
		{types.Stickers[4], types.NamazIndex[user.Language][4], taqvimTime.Isha},
	}

	maxLen := 0
	for _, e := range entries {
		l := utf8.RuneCountInString(e.Emoji + " " + e.Name)
		maxLen = max(maxLen, l)
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("📆 %s: %s, %s\n", messages.Messages[user.Language]["Today"], date, weekday))
	builder.WriteString(fmt.Sprintf("🏢 %s: %s\n\n", messages.Messages[user.Language]["Region"], region))
	builder.WriteString("<pre>\n")

	for _, e := range entries {
		builder.WriteString(fmt.Sprintf("%-8s %-7s: %s\n", e.Emoji, e.Name, e.Time))
	}

	builder.WriteString("</pre>")

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        builder.String(),
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(err)
		return
	}
}
