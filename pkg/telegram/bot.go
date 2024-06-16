package telegram

import (
	"financeBot/pkg/config"
	"financeBot/pkg/repository"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot        *tgbotapi.BotAPI
	postgresDb *repository.PostgresDB
	messages   *config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, postgresDb *repository.PostgresDB, messages *config.Messages) *Bot {
	return &Bot{bot: bot, postgresDb: postgresDb, messages: messages}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			if update.CallbackQuery != nil {
				b.handleCallback(update)
			} else {
				continue
			}
		} else if update.Message.IsCommand() {
			b.handlerCommands(update)
		} else {
			b.handleMessage(update)
		}
	}
}
