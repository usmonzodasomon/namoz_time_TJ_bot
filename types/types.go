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

var Stickers = []string{
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

//type NamazTimeSl struct {
//	Date  string
//	Namaz [5]NamazTimeStruct
//}

type NamazTime struct {
	Date        string `db:"date"`
	FajrFrom    string `db:"fajr_from"`
	FajrTo      string `db:"fajr_to"`
	ZuhrFrom    string `db:"zuhr_from"`
	ZuhrTo      string `db:"zuhr_to"`
	AsrFrom     string `db:"asr_from"`
	AsrTo       string `db:"asr_to"`
	MaghribFrom string `db:"maghrib_from"`
	MaghribTo   string `db:"maghrib_to"`
	IshaFrom    string `db:"isha_from"`
	IshaTo      string `db:"isha_to"`
}

type User struct {
	ChatID        int64  `db:"chat_id"`
	RegionID      int    `db:"region_id"`
	Username      string `db:"username"`
	Language      string `db:"lang"`
	LastMessageID int    `db:"last_message_id"`
	IsDeleted     bool   `db:"is_deleted"`
}

var SendNotifications map[int]map[int]bool = make(map[int]map[int]bool)

var RegionsID = map[string]int{
	"Душанбе":     1,
	"Истаравшан":  2,
	"Куляб":       3,
	"Худжанд":     4,
	"Рашт":        5,
	"Канибадам":   6,
	"Исфара":      7,
	"Ашт":         8,
	"Хорог":       9,
	"Мургаб":      10,
	"Кургантюбе":  11,
	"Пенджикент":  12,
	"Шахритус":    13,
	"Айни":        14,
	"Кӯлоб":       3,
	"Хуҷанд":      4,
	"Конибодом":   6,
	"Хоруғ":       9,
	"Мурғоб":      10,
	"Қурғонтеппа": 11,
	"Панҷакент":   12,
	"Шаҳритус":    13,
	"Aйнӣ":        14,
}

var Regions = map[string][]string{
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

var RegionsTime = map[int]int{
	1:  0,
	2:  -5,
	3:  -5,
	4:  -7,
	5:  -7,
	6:  -9,
	7:  -9,
	8:  -9,
	9:  -12,
	10: -20,
	11: 4,
	12: 5,
	13: 5,
	14: 5,
}

type Region struct {
	ID   int    `db:"id"`
	Name string `db:"region"`
}
