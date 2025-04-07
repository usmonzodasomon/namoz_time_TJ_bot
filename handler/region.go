package handler

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
	"log"
)

func (h *Handler) RegionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	kb := inline.New(b).Row()
	for i, region := range types.Regions[user.Language] {
		kb = kb.Button(region, []byte(region), h.onInlineKeyboardSelectRegion)
		if i%2 == 1 {
			kb = kb.Row()
		}
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        messages.Messages[user.Language]["ChooseRegion"] + ":",
		ReplyMarkup: &kb,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) onInlineKeyboardSelectRegion(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, data []byte) {
	if mes.Type == 1 {
		log.Println(fmt.Sprintf("MessageType is %s", "InaccessibleMessage"))
		return
	}

	user, err := h.storage.GetUser(mes.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	//b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
	//	CallbackQueryID: update.CallbackQuery.ID,
	//	ShowAlert:       false,
	//})

	//str := strings.TrimPrefix(update.CallbackQuery.Data, "region_")
	region := string(data)
	regionID := types.RegionsID[region]

	if err := h.storage.UpdateUser(types.User{
		ChatID:   mes.Message.Chat.ID,
		RegionID: regionID,
	}); err != nil {
		log.Println(err)
		return
	}

	//region, err := h.storage.GetRegionByID(user.Language, regionID)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Message.Chat.ID,
		Text:   messages.Messages[user.Language]["YourChoose"] + ": " + region,
	})
}
