package storage

import (
	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
)

type Storage interface {
	AddUserIfNotExist(user types.User) error
	GetUser(chatID int64) (types.User, error)
	GetAllUsersByRegionID(regionID int) ([]types.User, error)
	GetAllUsers() ([]types.User, error)
	UpdateUser(user types.User) error
	DeleteUser(chatID int64) error
	GetNamazTime(date string) (types.NamazTime, error)
	GetTaqvimTime() (types.TaqvimTime, error)
	GetRegionByID(lang string, id int) (types.Region, error)
	UpdateNamazTime(namazTime []types.NamazTime) error
	UpdateTaqvimTime(taqvimTime types.TaqvimTime) error
}
