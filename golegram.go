package golegram

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func NewBot(token string) (*Bot, error) {
	var bot Bot = Bot{Token: token}
	var err error

	bot.User, err = bot.getMe()

	if err == nil {
		return &bot, nil
	} else {
		return nil, err
	}
}

type msg struct {
	Chat_id                  string `json:"chat_id"`
	Text                     string `json:"text"`
	Parse_mode               string `json:"parse_mode"`
	Disable_web_page_preview bool   `json:"disable_web_page_preview"`
}

func (bot Bot) SendMessage(chat_id string, text string, disable_web_page_preview bool, parse_mode string) (message Message, err error) {
	var msg = msg{Chat_id: chat_id, Text: text, Disable_web_page_preview: disable_web_page_preview, Parse_mode: parse_mode}

	result, err := bot.sendCommand("sendMessage", msg)
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &message)

	return
}

type stickermsg struct {
	Chat_id string `json:"chat_id"`
	Sticker string `json:"sticker"`
}

func (bot Bot) SendSticker(chat_id string, fileId string) {
	var message Message

	var stickermsg = stickermsg{Chat_id: chat_id, Sticker: fileId}
	result, _ := bot.sendCommand("sendSticker", stickermsg)

	json.Unmarshal(result, message)
}

type fwd struct {
	Chat_id      string `json:"chat_id"`
	From_chat_id string `json:"from_chat_id"`
	Message_id   int32  `json:"message_id"`
}

func (bot Bot) ForwardMessage(chat_id string, from_chat_id string, message_id int32) (Message, error) {
	var message Message

	var fwd = fwd{Chat_id: chat_id, From_chat_id: from_chat_id, Message_id: message_id}
	result, err := bot.sendCommand("forwardMessage", fwd)
	if err != nil {
		return message, err
	}

	err1 := json.Unmarshal(result, &message)

	return message, err1
}

type getfile struct {
	File_id string `json:"file_id"`
}

func (bot Bot) GetFile(file_id string) (File, error) {
	var file File

	var req = getfile{File_id: file_id}
	result, err := bot.sendCommand("getFile", req)
	if err != nil {
		return file, err
	}

	err1 := json.Unmarshal(result, &file)

	return file, err1
}

type getupdate struct {
	Offset  int32 `json:"offset"`
	Limit   int32 `json:"limit"`
	Timeout int32 `json:"timeout"`
}

func (bot Bot) GetUpdates(offset int32, limit int32, timeout int32) ([]Update, error) {
	var update []Update

	var getupdate = getupdate{Offset: offset, Limit: limit, Timeout: timeout}
	result, err := bot.sendCommand("getUpdates", getupdate)
	if err != nil {
		return update, err
	}

	err1 := json.Unmarshal(result, &update)

	return update, err1
}

func StartWebhook(port int, cert string, key string) (error) {
	err := http.ListenAndServeTLS(":"+strconv.Itoa(port), cert, key, nil)

	return err
}

func (bot Bot) AddToWebhook(updatehandler UpdateHandler, pinghandler PingHandler){
	http.HandleFunc("/"+bot.Token, func(out http.ResponseWriter, in *http.Request) {
		handler(out, in, updatehandler)
	})

	http.HandleFunc("/"+bot.Token + "/ping", func(out http.ResponseWriter, in *http.Request) {
		pinghandler(out, in)
	})
}

func handler(out http.ResponseWriter, in *http.Request, updatehandler UpdateHandler) {
	var update Update
	json.NewDecoder(in.Body).Decode(&update)

	updatehandler(update)
}
