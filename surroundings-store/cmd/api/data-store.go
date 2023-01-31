package main

import (
	"fmt"
	"io/ioutil"
	dbmanage "main/db-manage"
	"net/http"
	"time"

	"github.com/tsawler/toolbox"
)

var tool toolbox.Tools

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	fmt.Println("hit the broker")
	_ = tool.WriteJSON(w, http.StatusOK, payload)
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

func (app *Config) DataStore(w http.ResponseWriter, r *http.Request) {
	tmp := &dbmanage.Surroundings{}
	err := tool.ReadJSON(w, r, tmp)
	tmp.Date = time.Now()
	fmt.Println(tmp)
	if err != nil {
		_ = tool.ErrorJSON(w, err)
	}
	err, client := dbmanage.DbConnection()
	if err != nil {
		_ = tool.ErrorJSON(w, err)
	}
	err = dbmanage.InsertSurroundings(client, *tmp)
	if err != nil {
		_ = tool.ErrorJSON(w, err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(nil)
}
