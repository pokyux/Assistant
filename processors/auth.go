package processors

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pokyux/Assistant/conf"
)

func IsNormalUser(_id int64) bool {
	for _, id := range conf.NormalUsers {
		if id == _id {
			return true
		}
	}

	return false
}

func WhoAmI(rcvd *tgbotapi.Message, rply *tgbotapi.MessageConfig) {
	rply.Text = fmt.Sprintf("Your ID is: %d, and you ", rcvd.From.ID)
	if rcvd.From.ID == conf.SuperAdmin {
		rply.Text += "are my admin."
	} else {
		rply.Text += "are not my admin."
	}
}

func AddNormalUser(rcvd *tgbotapi.Message, rply *tgbotapi.MessageConfig) {
	if rcvd.From.ID != conf.SuperAdmin {
		rply.Text = "Permission denied."
		return
	}

	textSplt := strings.Split(rcvd.Text, " ")
	if len(textSplt) != 2 {
		rply.Text = "Syntax error."
		return
	}

	id, err := strconv.ParseInt(textSplt[1], 10, 64)
	if err != nil {
		rply.Text = "Syntax error."
		return
	}

	conf.NormalUsers = append(conf.NormalUsers, id)
	rply.Text = fmt.Sprintf("User added. ID: %d.", id)
}
