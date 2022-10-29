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

func (app *Config)Speak(w http.ResponseWriter,r *http.Request){
	fmt.Println("reach speaker speak")
	var speakPayload SpeakPayload
	err := app.readJSON(w,r,&speakPayload)
	if err != nil{
		app.errorJSON(w,err)
		return
	}
	fmt.Println(r.Header)
	bytes,err := ioutil.ReadFile("sample/audio.wav")
	if err != nil{
		panic(err)
	}
	_,err = w.Write(bytes)
	if err != nil{
		panic(err)
	}
}