package processors

import "fmt"

func Repeat(rcvd *Message, rply *Reply) {
	rply.Text = fmt.Sprintf("Id: %d\nUsername: %s", rcvd.From.ID, rcvd.From.UserName)
	rply.Text += fmt.Sprintf("\nText: %s", rcvd.Text)
}
