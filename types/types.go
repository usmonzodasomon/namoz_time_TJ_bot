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

type TaqvimTime struct {
	Fajr    string `db:"fajr"`
	Zuhr    string `db:"zuhr"`
	Asr     string `db:"asr"`
	Maghrib string `db:"maghrib"`
	Isha    string `db:"isha"`
}

type User struct {
	ChatID           int64     `db:"chat_id"`
	RegionID         int       `db:"region_id"`
	Username         string    `db:"username"`
	Language         string    `db:"lang"`
	LastMessageID    int       `db:"last_message_id"`
	PrayerTimeSource string    `db:"prayer_time_source"`
	CreatedAt        time.Time `db:"created_at"`
	IsDeleted        bool      `db:"is_deleted"`
}

var SendNotifications map[int]map[int]bool = make(map[int]map[int]bool)

var RegionsID = map[string]int{
	"Душанбе":            1,
	"Истаравшан":         2,
	"Куляб":              3,
	"Худжанд":            4,
	"Рашт":               5,
	"Канибадам":          6,
	"Исфара":             7,
	"Ашт":                8,
	"Хорог":              9,
	"Мургаб":             10,
	"Кургантюбе":         11,
	"Пенджикент":         12,
	"Шахритус":           13,
	"Айни":               14,
	"Хамадони":           15,
	"Шамсиддин Шохин":    16,
	"Муминобод":          17,
	"Носири Хусрав":      18,
	"Турсунзода":         19,
	"Кӯлоб":              3,
	"Хуҷанд":             4,
	"Конибодом":          6,
	"Хоруғ":              9,
	"Мурғоб":             10,
	"Қурғонтеппа":        11,
	"Панҷакент":          12,
	"Шаҳритус":           13,
	"Aйнӣ":               14,
	"Ҳамадонӣ":           15,
	"Шамсиддин Шоҳин":    16,
	"Муъминобод":         17,
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
		"Хамадони",
		"Шамсиддин Шохин",
		"Муминобод",
		"Носири Хусрав",
		"Турсунзода",
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
		"Ҳамадонӣ",
		"Шамсиддин Шоҳин",
		"Муъминобод",
		"Носири Хусрав",
		"Турсунзода",
	},
}

var RegionsTime = map[int]int{
	1:  0,   // Душанбе
	2:  -1,  // Истаравшан
	3:  -4,  // Кӯлоб
	4:  -3,  // Хуҷанд
	5:  -6,  // Рашт
	6:  -6,  // Конибодом
	7:  -7,  // Исфара
	8:  -6,  // Ашт
	9:  -11, // Хоруғ
	10: -20, // Мурғоб
	11: 4,   // Қурғонтеппа
	12: 5,   // Панҷакент
	13: 3,   // Шаҳритус
	14: 1,   // Aйнӣ
	15: -3,  // Ҳамадонӣ
	16: -5,  // Шамсиддин Шоҳин
	17: -3,  // Муъминобод
	18: 4,   // Носири Хусрав
	19: 3,   // Турсунзода
}

type Region struct {
	ID   int    `db:"id"`
	Name string `db:"region"`
}

type UserStats struct {
	TotalUsers    int `db:"total_users"`
	ActiveUsers   int `db:"active_users"`
	NewUsersToday int `db:"new_users_today"`
}
