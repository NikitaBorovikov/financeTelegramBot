package main

import (
	"database/sql"
	"financeBot/pkg/config"
	"financeBot/pkg/repository"
	"financeBot/pkg/telegram"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)

	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	psglInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", cfg.Dbhost, cfg.Dbport, cfg.Dbuser, cfg.Password, cfg.Dbname)
	db, err := sql.Open("postgres", psglInfo)
	if err != nil {
		log.Printf("Connecting error")
		log.Fatal("Connecting error")
		db.Close()
	}
	log.Printf("successful conneting to DB")

	postgresDB := repository.NewPostgresDB(db)

	telegramBot := telegram.NewBot(bot, postgresDB, &cfg.Messages)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
