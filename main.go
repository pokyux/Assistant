package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/processors"
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
		bot.Send(Router(*update.Message))
	}
}

func Router(rcvd tgbotapi.Message) tgbotapi.MessageConfig {
	rply := tgbotapi.NewMessage(rcvd.Chat.ID, "")
	rply.ReplyToMessageID = rcvd.MessageID

	command := strings.Split(rcvd.Text, " ")[0]
	switch command {
	case "/oss":
		processors.UploadToOSS(&rcvd, &rply)
	default:
		processors.NotFound(&rcvd, &rply)
	}

	return rply
}
