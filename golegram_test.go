package golegram

import (
	"flag"
	"fmt"
	"testing"
)

var bot *Bot
var target string
var token string

func init() {
	flag.StringVar(&target, "target", "", "Chat_id as target")
	flag.StringVar(&token, "token", "", "Telegram Bot Token")
	flag.Parse()

	bot, _ = NewBot(token)
}
 
func TestSendMessage(t *testing.T) {

	if target == "" || token == "" {
		fmt.Println("Token and/or target is not set!")
		t.FailNow()
	}

	output, error := bot.SendMessage(target, "Testmessage")
	fmt.Println(output)
	if error != nil {
		fmt.Println(error.Error())
	}
}
