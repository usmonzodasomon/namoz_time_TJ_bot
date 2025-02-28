package handler

import (
	"log"
	"time"

	"github.com/go-telegram/bot/models"

	"echobot/messages"
	"echobot/types"
)

func (h *Handler) GetNamazTimeForCurrentRegion(namazTime types.NamazTime, regionID int) types.NamazTime {
	q := types.RegionsTime[regionID]
	FajrFrom, err := time.Parse("15:04", namazTime.FajrFrom)
	if err != nil {
		log.Println(err)
	}
	namazTime.FajrFrom = FajrFrom.Add(time.Minute * time.Duration(q)).Format("15:04")

	FajrTo, err := time.Parse("15:04", namazTime.FajrTo)
	if err != nil {
		log.Println(err)
	}
	namazTime.FajrTo = FajrTo.Add(time.Minute * time.Duration(q)).Format("15:04")

	ZuhrFrom, err := time.Parse("15:04", namazTime.ZuhrFrom)
	if err != nil {
		log.Println(err)
	}
	namazTime.ZuhrFrom = ZuhrFrom.Add(time.Minute * time.Duration(q)).Format("15:04")

	ZuhrTo, err := time.Parse("15:04", namazTime.ZuhrTo)
	if err != nil {
		log.Println(err)
	}
	namazTime.ZuhrTo = ZuhrTo.Add(time.Minute * time.Duration(q)).Format("15:04")

	AsrFrom, err := time.Parse("15:04", namazTime.AsrFrom)
	if err != nil {
		log.Println(err)
	}
	namazTime.AsrFrom = AsrFrom.Add(time.Minute * time.Duration(q)).Format("15:04")

	AsrTo, err := time.Parse("15:04", namazTime.AsrTo)
	if err != nil {
		log.Println(err)
	}
	namazTime.AsrTo = AsrTo.Add(time.Minute * time.Duration(q)).Format("15:04")

	MaghribFrom, err := time.Parse("15:04", namazTime.MaghribFrom)
	if err != nil {
		log.Println(err)
	}
	namazTime.MaghribFrom = MaghribFrom.Add(time.Minute * time.Duration(q)).Format("15:04")

	MaghribTo, err := time.Parse("15:04", namazTime.MaghribTo)
	if err != nil {
		log.Println(err)
	}
	namazTime.MaghribTo = MaghribTo.Add(time.Minute * time.Duration(q)).Format("15:04")

	IshaFrom, err := time.Parse("15:04", namazTime.IshaFrom)
	if err != nil {
		log.Println(err)
	}
	namazTime.IshaFrom = IshaFrom.Add(time.Minute * time.Duration(q)).Format("15:04")

	IshaTo, err := time.Parse("15:04", namazTime.IshaTo)
	if err != nil {
		log.Println(err)
	}
	namazTime.IshaTo = IshaTo.Add(time.Minute * time.Duration(q)).Format("15:04")

	return namazTime
}

func inlineButtonMain(lang string) *models.ReplyKeyboardMarkup {
	kb := &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: "🕓 " + messages.Messages[lang]["NamazTimeBtn"]},
			}, {
				{Text: "🇹🇯 " + messages.Messages[lang]["ChooseLanguageBtn"]},
				{Text: "🏙 " + messages.Messages[lang]["ChooseRegionBtn"]},
			},
			{
				{Text: "🕌 " + messages.Messages[lang]["Taqvim"]},
			},
		},
		ResizeKeyboard: true,
		Selective:      true,
	}
	return kb
}
