package handler

import (
	"log"
	"time"

	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/messages"
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
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
			},
			{
				{Text: "⚙️ " + messages.Messages[lang]["SettingsBtn"]},
				{Text: "🕌 " + messages.Messages[lang]["Taqvim"]},
			},
		},
		ResizeKeyboard: true,
		Selective:      true,
	}
	return kb
}

func (h *Handler) GetTaqvimTimeForCurrentRegion(taqvimTime types.TaqvimTime, regionID int) types.TaqvimTime {
	q := types.RegionsTime[regionID]

	Fajr, err := time.Parse("15:04", taqvimTime.Fajr)
	if err != nil {
		log.Println(err)
	} else {
		taqvimTime.Fajr = Fajr.Add(time.Minute * time.Duration(q)).Format("15:04")
	}

	Zuhr, err := time.Parse("15:04", taqvimTime.Zuhr)
	if err != nil {
		log.Println(err)
	} else {
		taqvimTime.Zuhr = Zuhr.Add(time.Minute * time.Duration(q)).Format("15:04")
	}

	Asr, err := time.Parse("15:04", taqvimTime.Asr)
	if err != nil {
		log.Println(err)
	} else {
		taqvimTime.Asr = Asr.Add(time.Minute * time.Duration(q)).Format("15:04")
	}

	Maghrib, err := time.Parse("15:04", taqvimTime.Maghrib)
	if err != nil {
		log.Println(err)
	} else {
		taqvimTime.Maghrib = Maghrib.Add(time.Minute * time.Duration(q)).Format("15:04")
	}

	Isha, err := time.Parse("15:04", taqvimTime.Isha)
	if err != nil {
		log.Println(err)
	} else {
		taqvimTime.Isha = Isha.Add(time.Minute * time.Duration(q)).Format("15:04")
	}

	return taqvimTime
}
