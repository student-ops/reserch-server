package main

import (
	"log"
	dbmanage "main/db-manage"
	"time"
)

func main() {
	// now := time.Now()
	// fmt.Println(now.)
	nowinfo := dbmanage.Surroundings{DeviceId: 4, Tempreture: 20.5, Airpressure: 1012, Date: time.Now()}
	err, client := dbmanage.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	dbmanage.InsertSurroundings(client, nowinfo)
}
