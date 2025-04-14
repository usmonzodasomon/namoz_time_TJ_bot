package handler

import (
	"context"
	"log"
	"sync"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/usmonzodasomon/namoz_time_TJ_bot/types"
)

func (h *Handler) MailingMeHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message.ReplyToMessage == nil {
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Ответьте на сообщение которое хотите отправить!",
			ReplyMarkup: inlineButtonMain(user.Language),
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

	_, err = b.CopyMessage(ctx, &bot.CopyMessageParams{
		ChatID:      update.Message.Chat.ID,
		FromChatID:  update.Message.Chat.ID,
		MessageID:   update.Message.ReplyToMessage.ID,
		ReplyMarkup: inlineButtonMain(user.Language),
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) MailingAllHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	user, err := h.storage.GetUser(update.Message.Chat.ID)
	if err != nil {
		log.Println(err)
		return
	}

	if update.Message.ReplyToMessage == nil {
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Ответьте на сообщение которое хотите отправить!",
			ReplyMarkup: inlineButtonMain(user.Language),
		})
		if err != nil {
			log.Println(err)
			return
		}
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Рассылка началась...",
		ReplyMarkup: inlineButtonMain(user.Language),
	})

	users, err := h.storage.GetAllUsers()
	if err != nil {
		log.Println(err)
		return
	}

	ch := make(chan types.User)
	wg := &sync.WaitGroup{}
	go func(ch chan types.User) {
		for _, u := range users {
			ch <- u
		}
		close(ch)
	}(ch)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for u := range ch {
				_, err = b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:      u.ChatID,
					Text:        update.Message.ReplyToMessage.Text,
					ReplyMarkup: inlineButtonMain(u.Language),
				})
				if err != nil {
					log.Println("error while sending mailing: " + err.Error())
				}
			}
		}()
	}
	wg.Done()

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Рассылка окончена...",
		ReplyMarkup: inlineButtonMain(user.Language),
	})
}
