package processors

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/conf"
)

func WhoAmI(rcvd *tgbotapi.Message, rply *tgbotapi.MessageConfig) {
	rply.Text = fmt.Sprintf("Your ID is: %d, and you ", rcvd.From.ID)
	if rcvd.From.ID == conf.SuperAdmin {
		rply.Text += "are my admin."
	} else {
		rply.Text += "are not my admin."
	}
}
