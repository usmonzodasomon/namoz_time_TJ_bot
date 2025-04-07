package handler

import (
	"github.com/usmonzodasomon/namoz_time_TJ_bot/storage"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}
