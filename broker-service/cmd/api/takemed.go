package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Config) Takemed(w http.ResponseWriter, r *http.Request, a TakemedPayload) error {

	jsonData, _ := json.MarshalIndent(a, "", "\t")
	// url := "http://medicineapi:8080/takemed"
	url := "http://medicine-api:3000/takemed"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	type Takemedresp struct {
		Message string
	}
	var tmr Takemedresp

	client := &http.Client{}
	response, err := client.Do(request)
	// err = tools.ReadJSON(w, request, &takemedMessage)
	json.NewDecoder(response.Body).Decode(&tmr)

	var speakpayload SpeakPayload
	speakpayload.Content = tmr.Message
	speakpayload.Speaker = 1
	fmt.Printf("takemd message :%s", tmr.Message)
	err = app.Speak(w, speakpayload)
	if err != nil {
		return err
	}
	return nil
}
