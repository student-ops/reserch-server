package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tsawler/toolbox"
)
var tools toolbox.Tools
type RequestPayload struct {
	Action string      `json:"action"`
	Speak  SpeakPayload `json:"voice,omitempty"`
}

type SpeakPayload struct{
	Speaker int `json:"speaker"`
	Content string `json:"content"`
}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := toolbox.JSONResponse{
		Error: false,
		Message: "Hit the broker",
	}
	fmt.Println("hit the broker")
	_=tools.WriteJSON(w,http.StatusOK,payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := tools.ReadJSON(w,r, &requestPayload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "echo":
		app.Echo(w)
	case "speak":
		fmt.Println("case speak")
		app.Speak(w,requestPayload.Speak)
	default:
		tools.ErrorJSON(w, errors.New("unknown action"))
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
func errhandle(err error,w http.ResponseWriter)bool{
	if err != nil{
		w.WriteHeader(400)
		w.Write(nil)
		return true
	}
	return false
}

func (app *Config) Speak(w http.ResponseWriter,a SpeakPayload){
	jsonData ,_ := json.MarshalIndent(a,"","\t")
	// url := "http://localhost:8080/speak"
	url := "http://speaker-service:8080/speak"
	request,err := http.NewRequest("POST",url,bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "mp3-binary")

	client := &http.Client{}
	response ,err := client.Do(request)
	if errhandle(err,w) {
		return
	}
	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		if errhandle(err,w) {
			return
		}
	} else if response.StatusCode != http.StatusAccepted {
		if errhandle(err,w) {
			return
		}
	}
	defer response.Body.Close()
	//it's need refacturing
	voice ,err:= ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println("voice read response error")
		w.WriteHeader(400)
		w.Write(nil)
	}
	fmt.Println(a)
	_,err = w.Write(voice)
}