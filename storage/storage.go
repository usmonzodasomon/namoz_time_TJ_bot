package storage

import (
	"echobot/types"
)

type Storage interface {
	AddUserIfNotExist(user types.User) error
	GetUser(chatID int64) (types.User, error)
	GetAllUsersByRegionID(regionID int) ([]types.User, error)
	UpdateUser(user types.User) error
	DeleteUser(chatID int64) error
	GetNamazTime(date string) (types.NamazTime, error)
	GetRegionByID(lang string, id int) (types.Region, error)
	UpdateNamazTime(namazTime []types.NamazTime) error
	//GetRegionID(chatID int64) (int, error)
	//GetLang(chatID int64) (string, error)
	//UpdateRegionID(chatID int64, regionID int) error
	//UpdateLanguage(chatID int64, language string) error
	//UpdateLastMessageID(chatID int64, lastMessageID int) error
	//GetLastMessageID(chatID int64) (int, error)
}
