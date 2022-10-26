package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)
type RequestPayload struct {
	Action string      `json:"action"`
	Voice  VoicePayload `json:"voice,omitempty"`
}

type VoicePayload struct{
	Speaker int `json:"speaker"`
	Content string `json:"content"`
}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "Hit the broker",
	}
	fmt.Println("hit the broker")
	app.writeJSON(w,http.StatusOK,payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "echo":
		app.Echo(w)
	default:
		app.errorJSON(w, errors.New("unknown action"))
		fmt.Println("unknown action")
	}
}


func (app *Config) Echo(w http.ResponseWriter) {
	bytes,err := ioutil.ReadFile("sample/audio.wav")
	if err != nil{
		panic(err)
	}
	state,err := w.Write(bytes)
	if err != nil{
		panic(err)
	}
	fmt.Printf("http respose state %d\n",state)
}
