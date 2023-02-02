package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TakemedPayload struct {
	UserId int `json:"userid"`
}
type SpeakPayload struct {
	Speaker int    `json:"speaker"`
	Content string `json:"content"`
}

func ExecRequest(a RequestPayload) error {
	var takemedpayload TakemedPayload
	takemedpayload.UserId = a.Id
	jsonData, _ := json.MarshalIndent(takemedpayload, "", "\t")
	takemed_request, err := http.NewRequest("POST", "http://mdedicine-api:3000", bytes.NewBuffer(jsonData))
	takemed_request.Header.Set("Content-Type", "mp3-binary")
	takemed_client := &http.Client{}
	response, err := takemed_client.Do(takemed_request)
	if err != nil {
		return err
	}
	type Takemedresp struct {
		Message string
	}
	var tmr Takemedresp
	json.NewDecoder(response.Body).Decode(&tmr)
	var speakpayload SpeakPayload
	speakpayload.Content = tmr.Message
	speakpayload.Speaker = 1
	jsonData, _ = json.MarshalIndent(speakpayload, "", "\t")
	speak_request, err := http.NewRequest("POST", "http://speaker-service:8080/speak", bytes.NewBuffer(jsonData))
	client := &http.Client{}
	response, err = client.Do(speak_request)
	voice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	url_device := "http://" + os.Getenv("ADDRESS") + ":8000/speak"
	request_device, err := http.NewRequest("POST", url_device, bytes.NewBuffer(voice))
	if err != nil {
		return err
	}
	response, err = client.Do(request_device)
	if err != nil {
		return err
	}
	fmt.Println(response)
	defer response.Body.Close()
	return nil
}
