package bot

import (
	"time"

	"github.com/Impisigmatus/service_core/log"
	"github.com/Impisigmatus/service_core/telegram"
	"github.com/jmoiron/sqlx"

	tg_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	botTG *telegram.Telegram
	db    *sqlx.DB
}

type Reciever struct {
	duration time.Duration
	offset   int
	handle   *tg_bot.MessageConfig
}

func NewBot(token string, db *sqlx.DB) *Bot {
	log.Init(log.LevelInfo)
	log.Info("Hui")
	return &Bot{
		botTG: telegram.New(token, &Reciever{
			duration: 5 * time.Second,
			offset:   0,
			handle:   nil,
		}),
		db: db,
	}
}

func (reciever *Reciever) GetOffset() int {
	return reciever.offset
}

func (reciever *Reciever) GetTimeout() time.Duration {
	return reciever.duration
}

func (reciever *Reciever) Handle(payload *tg_bot.Message) (*tg_bot.MessageConfig, error) {
	log.Info(payload.Text)

	reply := tg_bot.NewMessage(payload.Chat.ID, payload.Text)

	return &reply, nil
}
