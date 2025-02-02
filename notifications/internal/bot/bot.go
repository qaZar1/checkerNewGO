package bot

import (
	"time"

	"github.com/Impisigmatus/service_core/log"
	"github.com/Impisigmatus/service_core/telegram"
	"github.com/qaZar1/checkerNewGO/notifications/internal/api"
	"github.com/qaZar1/checkerNewGO/notifications/internal/models"
	"github.com/sirupsen/logrus"

	tg_bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	botTG *telegram.Telegram
}

type Reciever struct {
	duration time.Duration
	offset   int
	api      *api.APIUsers
}

func NewBot(token string, api *api.APIUsers) *Bot {
	log.Init(log.LevelInfo)
	return &Bot{
		botTG: telegram.New(token, &Reciever{
			duration: 5 * time.Second,
			offset:   0,
			api:      api,
		}),
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

	ok, err := reciever.api.GetUserByChatID(payload.From.ID)
	if err != nil {
		logrus.Errorf("Invalid get user by chat id: %s", err)
	}
	if !ok {
		reply, err := reciever.authorization(payload)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}

	return nil, nil
}

func (reciever *Reciever) authorization(payload *tg_bot.Message) (*tg_bot.MessageConfig, error) {
	ok, err := reciever.api.AddUser(models.User{
		Chat_ID:  payload.From.ID,
		Username: payload.From.UserName,
		Name:     payload.From.FirstName,
	})
	if err != nil {
		logrus.Errorf("Invalid add user: %s", err)
	}

	if ok {
		reply := tg_bot.NewMessage(payload.Chat.ID, "Вы успешно зарегистрировались в боте. Теперь вам будут приходить уведомления о появлении новых версий языка Go.")
		return &reply, nil
	} else {
		reply := tg_bot.NewMessage(payload.Chat.ID, "Произошла ошибка при регистрации")
		return &reply, err
	}
}
