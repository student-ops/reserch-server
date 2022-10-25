package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"io"
)
type RequestPayload struct {
	Action string      `json:"action"`
	Voice   VoicePayload `json:"voice,omitempty"`
}


func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error: false,
		Message: "Hit the broker",
	}
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
	case "voice":
		app.voice(w, requestPayload.Voice)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) Voice(w http.ResponseWriter, a VoicePayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://speaker-service/speak:8080", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, response.Body); err != nil {
		app.errorJSON(w,errors.New("error copying respose from speaker"))
		return
	}

	// err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}


	app.writeJSON(w, http.StatusAccepted, payload)
}

