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

type fwd struct{
	Chat_id int32 `json:"chat_id"`
	From_chat_id int32 `json:"from_chat_id"`
	Message_id int32 `json:"message_id"`
}

func (bot Bot) ForwardMessage(chat_id int32, from_chat_id int32, message_id int32) (Message, error) {
	var message Message

	var fwd = fwd{ Chat_id: chat_id, From_chat_id: from_chat_id, Message_id: message_id,}
	result, err := bot.sendCommand("forwardMessage", fwd)
	if err != nil {
		return message, err
	}

	err1 := json.Unmarshal(result, &message)

	return message, err1
}

type getfile struct{
	File_id string `json:"file_id"`
}

func (bot Bot) GetFile(file_id string) (File, error) {
	var file File

	var req = getfile{ File_id: file_id,}
	result, err := bot.sendCommand("getFile", req)
	if err != nil {
		return file, err
	}

	err1 := json.Unmarshal(result, &file)

	return file, err1
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