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
	Action       string               `json:"action"`
	Speak        SpeakPayload         `json:"speak,omitempty"`
	Surroundings SurroundingsPalyload `json:"surroundings,omitempty`
	Takemed      TakemedPayload       `json:"takemed,omitempty`
}

type SpeakPayload struct {
	Speaker int    `json:"speaker"`
	Content string `json:"content"`
}

type SurroundingsPalyload struct {
	Tempreture int `json:"tempreture"`
}

type TakemedPayload struct {
	UserId int `json:"userid"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	fmt.Println("hit the broker")
	_ = tools.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "echo":
		app.Echo(w)
	case "speak":
		fmt.Println("case speak")
		err = app.Speak(w, requestPayload.Speak)
		if err != nil {
			http.Error(w, "speak error", 401)
		}
	case "surroundings":
		fmt.Println("case surroundings")
		err = app.SurroundingsStore(w, requestPayload.Surroundings)
		if err != nil {
			http.Error(w, "surroudings error", 402)
		}

	case "takemed":
		fmt.Println("case takemed")
		err = app.Takemed(w, requestPayload.Takemed)
		if err != nil {
			http.Error(w, "takemed error")
		}
	default:
		tools.ErrorJSON(w, errors.New("unknown action"))
		fmt.Println("unknown action")
	}
}

func (app *Config) Echo(w http.ResponseWriter) {
	bytes, err := ioutil.ReadFile("sample/audio.wav")
	if err != nil {
		panic(err)
	}
	state, err := w.Write(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("http respose state %d\n", state)
}
func errhandle(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.WriteHeader(400)
		w.Write(nil)
		return true
	}
	return false
}
func (app *Config) SurroundingsStore(w http.ResponseWriter, a SurroundingsPalyload) error {
	return nil
}

// needs erorr handling
func (app *Config) Speak(w http.ResponseWriter, a SpeakPayload) error {
	jsonData, _ := json.MarshalIndent(a, "", "\t")
	// url := "http://localhost:8080/speak"
	url := "http://speaker-service:8080/speak"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "mp3-binary")

	client := &http.Client{}
	response, err := client.Do(request)
	if errhandle(err, w) {
		return err
	}
	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		return err
	} else if response.StatusCode != http.StatusAccepted {
		return nil
	}
	//it's need refacturing	defer response.Body.Close()
	voice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(a)
	_, err = w.Write(voice)
	if err != nil {
		return err
	}
	return nil
}
