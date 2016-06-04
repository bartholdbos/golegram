package golegram

import(
	"net/http"
	"encoding/json"
	"io"
	"errors"
)

type response struct{
	Ok				bool
	Result			json.RawMessage
	Error_code		int32
	Description		string
}

func (bot Bot) sendCommand(method string, params interface{}) (json.RawMessage, error){
	reader, writer := io.Pipe()
	var err error

	go func() (){
		defer writer.Close()

		err = json.NewEncoder(writer).Encode(&params)
	}()

	resp, err1 := http.Post("https://api.telegram.org/bot" + bot.Token + "/" + method, "application/json", reader)
	if err != nil {
		return nil, err
	}
	if err1 != nil {
		return nil, err1
	}

	defer resp.Body.Close()

	var r response //change
	err2 := json.NewDecoder(resp.Body).Decode(&r);

	if err2 != nil {
		return nil, err2
	}

	if(!r.Ok){
		return nil, errors.New(r.Description)
	}

	return r.Result, nil
}

func (bot Bot) getMe() (User, error) {
	var user User

	result, err := bot.sendCommand("getMe", nil)
	if err != nil {
		return user, err
	}

	err1 := json.Unmarshal(result, &user)

	return user, err1
}