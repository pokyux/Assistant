package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/conf"
	"github.com/pokyux/Assistant/processors"
)

var router map[string]func(*tgbotapi.Message, *tgbotapi.MessageConfig)

func main() {
	processors.InitOSS()
	InitRouter()

	bot, err := tgbotapi.NewBotAPI(conf.TGBotToken)
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

func InitRouter() {
	router = make(map[string]func(*tgbotapi.Message, *tgbotapi.MessageConfig))
	router["/oss"] = processors.UploadToOSS
	router["/whoami"] = processors.WhoAmI
}

func Router(rcvd tgbotapi.Message) tgbotapi.MessageConfig {
	rply := tgbotapi.NewMessage(rcvd.Chat.ID, "")
	rply.ReplyToMessageID = rcvd.MessageID

	processor := router[strings.Split(rcvd.Text, " ")[0]]
	if processor == nil {
		processor = processors.NotFound
	}

	processor(&rcvd, &rply)
	return rply
}
