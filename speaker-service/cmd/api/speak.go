package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type SpeakPayload struct{
	Speaker int `json:"speaker"`
	Content string `json:"content"`
}

func (app *Config)speak(w http.ResponseWriter,r *http.Request){
	var speakPayload SpeakPayload
	err := app.readJSON(w,r,&speakPayload)
	if err != nil{
		app.errorJSON(w,err)
		return
	}
	fmt.Println(speakPayload)
	bytes,err := ioutil.ReadFile("sample/audio.wav")
	if err != nil{
		panic(err)
	}
	_,err = w.Write(bytes)
	if err != nil{
		panic(err)
	}
}