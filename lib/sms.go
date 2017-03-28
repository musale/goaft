package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// AfricastalkingGateway struct
type AfricastalkingGateway struct {
	Username string
	APIKEY   string
	Debug    bool
	Format   string
}

// AfricastalkigGatewayInterface with methods
type AfricastalkigGatewayInterface interface {
	SendMessage() []APIRecipient
}

// APIRecipient - number receiving SMS
type APIRecipient struct {
	Number    string  `json:"number"`
	Status    string  `json:"status"`
	Cost      float64 `json:"cost"`
	MessageID string  `json:"messageId"`
}

// ErrorDetail - number receiving SMS
type ErrorDetail struct {
	Source string `json:"source"`
	Status string `json:"status"`
	Error  error
}

// MessageData - with msg & list of recipients
type MessageData struct {
	Message    string         `json:"message"`
	Recipients []APIRecipient `json:"recipients"`
}

// SendMessage method
func (gateway *AfricastalkingGateway) SendMessage(to, message, senderID string) (*ErrorDetail, []APIRecipient) {
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
		return &ErrorDetail{"CREATE SEND SMS REQ", "Failed", err}, []APIRecipient{}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Add("Accept", "Application/json")
	req.Header.Add("apikey", gateway.APIKEY)

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Do Error: ", err)
		return &ErrorDetail{"POST SEND SMS REQ", "Failed", err}, []APIRecipient{}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("readall error: ", err)
		return &ErrorDetail{"READ SEND SMS RES", "Failed", err}, []APIRecipient{}
	}

	retData := make(map[string]MessageData)

	err = json.Unmarshal(body, &retData)

	if err != nil {
		log.Println("Marshal Error: ", err)
		return &ErrorDetail{"MARSHAL SEND SMS RES", "Failed", err}, []APIRecipient{}
	}

	return nil, retData["SMSMessageData"].Recipients
}

func getSMSUrl() string {
	return "http://127.0.0.1:4027/aft"
	// return "http://africastalking.com"
}
