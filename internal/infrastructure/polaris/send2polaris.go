package polaris

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Messages []struct {
	ChatId  string `json:"chatid"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ReceivedMessage struct {
	Content string `json:"content"`
}

func ReplyFromSlack(fromSlack string, PayloadTS string) string {
	url := "https://sure-cheaply-kite.ngrok-free.app/entries"
	payload := Messages{
		{
			ChatId:  PayloadTS,
			Role:    "user",
			Content: fromSlack,
		},
	}
	aiResponse, _ := makePostRequest(url, payload)
	return aiResponse.Content
}

func makePostRequest(url string, payload Messages) (ReceivedMessage, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return ReceivedMessage{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return ReceivedMessage{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ReceivedMessage{}, err
	}

	var receivedMessage ReceivedMessage
	err = json.Unmarshal(body, &receivedMessage)
	if err != nil {
		return ReceivedMessage{}, err
	}

	return receivedMessage, nil
}
