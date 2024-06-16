package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errAddData   = errors.New("error add to db")
	errDelData   = errors.New("error delete data from db ")
	errEnterData = errors.New("error enter data")
)

func (b *Bot) handleErrors(user_id int64, err error) {
	switch err {
	case errAddData:
		text := b.messages.ErrorAddDataMessage
		msg := tgbotapi.NewMessage(user_id, text)
		b.bot.Send(msg)
	case errDelData:
		text := b.messages.ErrorDelDataMessage
		msg := tgbotapi.NewMessage(user_id, text)
		b.bot.Send(msg)
	case errEnterData:
		text := b.messages.ErrorEnterData
		msg := tgbotapi.NewMessage(user_id, text)
		b.bot.Send(msg)
	default:
		text := "Ошибка"
		msg := tgbotapi.NewMessage(user_id, text)
		b.bot.Send(msg)
	}
}
