package goaft

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type AfricastalkingGateway struct {
	Username string
	APIKEY   string
	Dest     string
}

type AfricastalkigGatewayInterface interface {
	SendMessage() []APIRecipient
}

type APIRecipient struct {
	Number    string  `json:"number"`
	Status    string  `json:"status"`
	Cost      float64 `json:"cost"`
	MessageId string  `json:"messageId"`
}

type MessageData struct {
	Message    string         `json:"message"`
	Recipients []APIRecipient `json:"recipients"`
}

func (gateway *AfricastalkingGateway) SendMessage(to, message, senderID string) []APIRecipient {
	client := http.Client{}
	form := url.Values{}
	form.Add("username", gateway.Username)
	form.Add("message", message)
	form.Add("to", to)
	form.Add("from", senderID)
	req, err := http.NewRequest(
		"POST", getSMSUrl(), strings.NewReader(form.Encode()))

	if err != nil {
		log.Println("Request Error: ", err)
		return []APIRecipient{}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("Accept", "Application/json")
	req.Header.Add("apikey", gateway.APIKEY)

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Do Error: ", err)
		return []APIRecipient{}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("readall error: ", err)
		return []APIRecipient{}
	}

	retData := make(map[string]MessageData)

	err = json.Unmarshal(body, &retData)

	if err != nil {
		log.Println("Marshal Error: ", err)
		return []APIRecipient{}
	}

	return retData["SMSMessageData"].Recipients
}

func getSMSUrl() string {
	return "http://127.0.0.1:4027/aft"
	// return "http://africastalking.com"
}
