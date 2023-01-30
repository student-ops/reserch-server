package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/tsawler/toolbox"
)

func main() {

	type takemdpayload struct {
		Userid int `json:"userid"`
	}
	type takemdres struct {
		Message string `json:"message"`
	}
	var jsondata takemdpayload
	jsondata.Userid = 1
	jsonData, _ := json.MarshalIndent(jsondata, "", "\t")
	// url := "http://medicineapi:8080/takemed"
	url := "http://localhost:3000/takemed"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	var tool = toolbox.Tools
	tool.ReadJSON(w, request, &res)

	// err = tools.ReadJSON(w, request, &takemedMessage)
	if err != nil {
		log.Panic(err)
	}

	// var speakpayload SpeakPayload
	// speakpayload.Content = takemedMessage.message
	// speakpayload.Speaker = 1
	// fmt.Printf("takemd message :%s", takemedMessage)
	// err = app.Speak(w, speakpayload)
	// if err != nil {
	// return err
	// }
}
