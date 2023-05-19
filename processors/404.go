package processors

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func NotFound(rcvd *tgbotapi.Message, rply *tgbotapi.MessageConfig) {
	rply.Text = "Command not found!"
}
