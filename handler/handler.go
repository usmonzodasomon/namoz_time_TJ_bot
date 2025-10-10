package handler

import (
	"github.com/usmonzodasomon/namoz_time_TJ_bot/storage"
)

type Handler struct {
	storage     storage.Storage
	adminChatID int64
}

func NewHandler(storage storage.Storage, adminChatID int64) *Handler {
	return &Handler{
		storage:     storage,
		adminChatID: adminChatID,
	}
}
