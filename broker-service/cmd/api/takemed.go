package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (app *Config) Takemed(w http.ResponseWriter, a TakemedPayload) error {
	jsonData, _ := json.MarshalIndent(a, "", "\t")
	url := "http://medicineapi:8080/takemed"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application-json")
	type TakemdMessage struct {
		message string
	}
	var takemedMessage TakemdMessage
	err = tools.ReadJSON(w, request, &takemedMessage)
	if err != nil {
		return err
	}
	var speakpayload SpeakPayload
	speakpayload.Content = takemedMessage.message
	speakpayload.Speaker = 1
	err = app.Speak(w, speakpayload)
	if err != nil {
		return err
	}
	return nil
}
