package types

import "time"

var NamazIndex map[string]map[int]string = map[string]map[int]string{
	"ru": {
		0: "Фаджр",
		1: "Зухр",
		2: "Аср",
		3: "Магриб",
		4: "Иша",
	},
	"tj": {
		0: "Бомдод",
		1: "Пешин",
		2: "Аср",
		3: "Шом",
		4: "Хуфтан",
	},
}

var Stickers []string = []string{
	"🏙",
	"🌅",
	"🌇",
	"🌌",
	"🌃",
}

type StringNamazTime struct {
	Day     string
	Today   string
	Fajr    string
	Dhuhr   string
	Asr     string
	Maghrib string
	Isha    string
}

type NamazTimeStruct struct {
	From time.Time
	To   time.Time
}

type NamazTime struct {
	Today time.Time
	Namaz [5]NamazTimeStruct
}

type User struct {
	ChatID   int64
	RegionID int64
	Username string
	Language string
}

var SentNotifications map[int]map[int]bool = make(map[int]map[int]bool)

var RegionsID map[string]int = map[string]int{
	"Душанбе":     0,
	"Истаравшан":  1,
	"Куляб":       2,
	"Худжанд":     3,
	"Рашт":        4,
	"Канибадам":   5,
	"Исфара":      6,
	"Ашт":         7,
	"Хорог":       8,
	"Мургаб":      9,
	"Кургантюбе":  10,
	"Пенджикент":  11,
	"Шахритус":    12,
	"Айни":        13,
	"Кӯлоб":       2,
	"Хуҷанд":      3,
	"Конибодом":   5,
	"Хоруғ":       8,
	"Мурғоб":      9,
	"Қурғонтеппа": 10,
	"Панҷакент":   11,
	"Шаҳритус":    12,
	"Aйнӣ":        13,
}

var Regions map[string][]string = map[string][]string{
	"ru": {
		"Душанбе",
		"Истаравшан",
		"Куляб",
		"Худжанд",
		"Рашт",
		"Канибадам",
		"Исфара",
		"Ашт",
		"Хорог",
		"Мургаб",
		"Кургантюбе",
		"Пенджикент",
		"Шахритус",
		"Айни",
	},
	"tj": {
		"Душанбе",
		"Истаравшан",
		"Кӯлоб",
		"Хуҷанд",
		"Рашт",
		"Конибодам",
		"Исфара",
		"Ашт",
		"Хоруғ",
		"Мурғоб",
		"Қурғонтеппа",
		"Панҷакент",
		"Шаҳритус",
		"Aйнӣ",
	},
}

var RegionsTime map[int]int = map[int]int{
	0:  0,
	1:  -5,
	2:  -5,
	3:  -7,
	4:  -7,
	5:  -9,
	6:  -9,
	7:  -9,
	8:  -12,
	9:  -20,
	10: 4,
	11: 5,
	12: 5,
	13: 5,
}
