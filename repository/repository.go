package repository

import (
	"echobot/types"
)

type Repository interface {
	CreateUser(user types.User) error
	GetRegionID(chatID int64) (int, error)
	GetAllUsersByRegionID(regionID int) ([]int64, error)
	GetLang(chatID int64) (string, error)
	UpdateRegionID(chatID int64, regionID int) error
	UpdateLanguage(chatID int64, language string) error
	UpdateLastMessageID(chatID int64, lastMessageID int) error
	GetLastMessageID(chatID int64) (int, error)
	DeleteUser(chatID int64) error
}
