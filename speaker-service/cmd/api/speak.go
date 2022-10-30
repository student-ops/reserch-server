package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"

	//mod fileからのパスで指定
	"main/cmd/api/data"
	"net/http"
	"strconv"

	"github.com/tsawler/toolbox"
)

type SpeakPayload struct {
	Speaker int    `json:"speaker"`
	Content string `json:"content"`
}

var tool toolbox.Tools

type config struct {
	endpoint   string
	speaker    int
	style      int
	speed      float64
	intonation float64
	volume     float64
	pitch      float64
	output     string
}

func (app *Config) Speak(w http.ResponseWriter, r *http.Request) {
	speakPayload := &SpeakPayload{}
	err := tool.ReadJSON(w, r, speakPayload)
	if err != nil {
		_ = tool.ErrorJSON(w, err)
		return
	}
	b := app.byteRecive(speakPayload)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Header)
	fmt.Println(speakPayload.Content)
	_, err = w.Write(b)
	if err != nil {
		panic(err)
	}
}

func getQuery(conf config, id int, text string) (*data.Params, error) {
	req, err := http.NewRequest("POST", conf.endpoint+"/audio_query", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("speaker", strconv.Itoa(id))
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()
	//log.Println(req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var params *data.Params
	if err := json.NewDecoder(resp.Body).Decode(&params); err != nil {
		return nil, err
	}
	return params, nil
}

func synth(url string, id int, params *data.Params) ([]byte, error) {
	b, err := json.MarshalIndent(params, "", "  ")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url+"/synthesis", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "audio/wav")
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("speaker", strconv.Itoa(id))
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, resp.Body); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
func (app *Config) byteRecive(speakPayload *SpeakPayload) []byte {
	var conf config
	conf.endpoint = "http://localhost:50021"
	speaker := speakPayload.Speaker
	text := speakPayload.Content
	fmt.Printf("byteRecive text %s", text)
	params, err := getQuery(conf, speaker, text)
	if err != nil {
		log.Fatal(err)
	}
	params.SpeedScale = 1.0
	params.PitchScale = 0.0
	params.IntonationScale = 1.0
	params.VolumeScale = 1.0

	b, err := synth(conf.endpoint, speaker, params)
	//send messagec
	return b
}
