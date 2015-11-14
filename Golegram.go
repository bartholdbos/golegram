package Golegram

import(
	"encoding/json"
)

func NewBot(token string) (*Bot, error){
	var bot Bot = Bot{ Token: token, }
	var err error

	bot.User, err = bot.getMe()

	if err == nil {
		return &bot, nil
	}else{
		return nil, err
	}
}

type msg struct{
	Chat_id int32 `json:"chat_id"`
	Text string `json:"text"`
}

func (bot Bot) SendMessage(chat_id int32, text string) (Message, error) {
	var message Message

	var msg = msg{ Chat_id: chat_id, Text: text,}
	result, err := bot.sendCommand("sendMessage", msg)
	if err != nil {
		return message, err
	}

	err1 := json.Unmarshal(result, &message)

	return message, err1
}

type getupdate struct{
	Offset int32 `json:"offset"`
	Limit int32 `json:"limit"`
	Timeout int32 `json:"timeout"`
}

func (bot Bot) GetUpdates(offset int32, limit int32, timeout int32) ([]Update, error){
	var update []Update

	var getupdate = getupdate{Offset: offset, Limit: limit, Timeout: timeout,}
	result, err := bot.sendCommand("getUpdates", getupdate)
	if err != nil {
		return update, err
	}

	err1 := json.Unmarshal(result, &update)

	return update, err1
}