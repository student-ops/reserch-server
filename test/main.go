package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:3000/takemed?id=1", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b))

}
