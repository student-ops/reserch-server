package main

import (
	"fmt"
	"log"
	dbmanage "main/db-manage"
	"math/rand"
	"time"
)

func main() {
	err, client := dbmanage.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < 180; i++ {
		now := time.Now()
		now = now.Add(-time.Minute * time.Duration(i))
		temp := 10 + 20*rand.Float32()
		airp := 1000 + 20*rand.Float32()
		fmt.Println(now)
		d := dbmanage.Surroundings{1, temp, airp, now}
		dbmanage.InsertSurroundings(client, d)
	}
}
