package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// func main() {
// 	data := []byte("this is some data stored as a byte slice in Go Lang!")

// 	// convert byte slice to io.Reader
// 	reader := bytes.NewReader(data)

// 	// read only 4 byte from our io.Reader
// 	buf := make([]byte, 4)
// 	n, err := reader.Read(buf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(string(buf[:n]))
// }

func main() {
	byte, err := ioutil.ReadFile("sample.wav")
	if err != nil {
		fmt.Println(29)
		fmt.Println(err)
	}

	url := "http://localhost:8000/foo"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(byte))
	request.Header.Set("Content-Type", "mp3-binary")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	response, _ := client.Do(request)
	if err != nil {
	}
	defer response.Body.Close()
}
