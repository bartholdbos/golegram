package golegram

import (
	"flag"
	"fmt"
	"testing"
)

var bot *Bot
var target int
var token string

func init() {
	flag.IntVar(&target, "target", 0, "Chat_id as target")
	flag.StringVar(&token, "token", "", "Telegram Bot Token")
	flag.Parse()

	bot, _ = NewBot(token)
}

func TestSendMessage(t *testing.T) {

	if target == 0 || token == "" {
		fmt.Println("Token and/or target is not set!")
		t.FailNow()
	}

	output, error := bot.SendMessage(int32(target), "Testmessage")
	fmt.Println(output)
	if error != nil {
		fmt.Println(error.Error())
	}
}
