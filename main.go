package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Args[1]
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	log.Println("Bot created.")

	uConfig := tgbotapi.NewUpdate(0)
	uConfig.Timeout = 60
	updates := bot.GetUpdatesChan(uConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("Msg from: %s.\n", update.Message.From.UserName)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "rcvd")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
