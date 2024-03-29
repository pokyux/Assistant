package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/conf"
	"github.com/pokyux/Assistant/global"
	"github.com/pokyux/Assistant/processors"
)

var router map[string]func(*tgbotapi.Message, *tgbotapi.MessageConfig)

func main() {
	InitFlags()
	processors.InitOSS()
	InitRouter()

	var err error
	global.Bot, err = tgbotapi.NewBotAPI(conf.TGBotToken)
	if err != nil {
		panic(err)
	}
	log.Println("Bot created.")

	uConfig := tgbotapi.NewUpdate(0)
	uConfig.Timeout = 60
	updates := global.Bot.GetUpdatesChan(uConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if global.Debug {
			log.Printf("uid: %d, Msg: %s\n", update.Message.From.ID, update.Message.Text)
		}
		global.Bot.Send(Router(*update.Message))
	}
}

func InitRouter() {
	router = make(map[string]func(*tgbotapi.Message, *tgbotapi.MessageConfig))
	router["oss"] = processors.UploadToOSS
	router["whoami"] = processors.WhoAmI
	router["addnormaluser"] = processors.AddNormalUser
	router["repeat"] = processors.Repeat
}

func Router(rcvd tgbotapi.Message) tgbotapi.MessageConfig {
	rply := tgbotapi.NewMessage(rcvd.Chat.ID, "")
	rply.ReplyToMessageID = rcvd.MessageID

	processor := router[rcvd.Command()]

	if rcvd.Document != nil {
		processor = processors.UploadToOSS
	}

	if processor == nil {
		processor = processors.NotFound
	}

	processor(&rcvd, &rply)
	return rply
}

func InitFlags() {
	global.Debug = false

	for _, flag := range os.Args {
		switch flag {
		case "--debug":
			log.Println("Debug mode: on")
			global.Debug = true
		}
	}
}
