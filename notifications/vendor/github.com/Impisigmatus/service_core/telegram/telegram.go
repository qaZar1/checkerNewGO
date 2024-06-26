package telegram

import (
	"fmt"
	"time"

	"github.com/Impisigmatus/service_core/log"
	tg_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//go:generate mockgen -source=telegram.go -package mocks -destination ../autogen/mocks/telegram.go
type IReciever interface {
	Handle(payload *tg_bot.Message) (*tg_bot.MessageConfig, error)
	GetOffset() int
	GetTimeout() time.Duration
}

type _ITelegramAPI interface {
	Send(c tg_bot.Chattable) (tg_bot.Message, error)
	GetUpdatesChan(cfg tg_bot.UpdateConfig) tg_bot.UpdatesChannel
}

type Telegram struct {
	api      _ITelegramAPI
	reciever IReciever
}

func New(token string, reciever IReciever) *Telegram {
	if reciever == nil {
		log.Panicf("Invalid init: reciever is nil")
	}
	api, err := tg_bot.NewBotAPI(token)
	if err != nil {
		log.Panicf("Invalid telegram api: %s", err)
	}

	tg := newTelegram(api, reciever)
	tg.consume()
	return tg
}

func newTelegram(api _ITelegramAPI, reciever IReciever) *Telegram {
	return &Telegram{api: api, reciever: reciever}
}

func (tg *Telegram) Send(chatID uint64, data string) error {
	msg := tg_bot.NewMessage(int64(chatID), data)
	if _, err := tg.api.Send(msg); err != nil {
		return fmt.Errorf("Invalid send: %s", err)
	}

	return nil
}

func (tg *Telegram) consume() {
	go func() {
		updater := tg_bot.NewUpdate(tg.reciever.GetOffset())
		updater.Timeout = int(tg.reciever.GetTimeout().Seconds())
		updates := tg.api.GetUpdatesChan(updater)

		for update := range updates {
			if update.Message != nil {
				msg, err := tg.reciever.Handle(update.Message)
				if err != nil {
					log.Error("Invalid handle msg", err)
					continue
				}
				if _, err := tg.api.Send(msg); err != nil {
					log.Error("Invalid send", err)
					continue
				}
			}
		}
	}()
}
