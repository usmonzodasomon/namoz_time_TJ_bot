package handler

import (
	"github.com/usmonzodasomon/namoz_time_TJ_bot/storage"
)

type SchedulerInterface interface {
	SendDatabaseBackup()
}

type Handler struct {
	storage     storage.Storage
	adminChatID int64
	scheduler   SchedulerInterface
}

func NewHandler(storage storage.Storage, adminChatID int64) *Handler {
	return &Handler{
		storage:     storage,
		adminChatID: adminChatID,
	}
}

func (h *Handler) SetScheduler(scheduler SchedulerInterface) {
	h.scheduler = scheduler
}
