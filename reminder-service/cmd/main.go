package main

import (
	"fmt"
)

// func main() {
// 	now := time.Now()
// 	_, offset := now.Zone()
// 	now = now.Add(time.Duration(32400 - offset))
// 	fmt.Println(now)

// }

type RequestPayload struct {
	Id   int `json:"id"`
	time int `json:"time"`
}

func main() {
	var requestpayload RequestPayload
	requestpayload.Id = 1
	requestpayload.time = 1
	// for true {
	// 	now := time.Now()
	// 	if now.Second() == 0 {
	// 		break
	// 	}
	// 	time.Sleep(10)
	// }
	// for true {
	// 	now := time.Now()
	// 	if (now.Minute()) == 50 {
	// 		break
	// 	}
	// 	time.Sleep(1 * time.Minute)
	// }
	// for i := 0; i < 24; i++ {
	// 	now := time.Now()
	// 	_, offset := now.Zone()
	// 	now = now.Add(time.Duration(32400 - offset))
	// 	if now.Hour() == 8 {
	// 		requestpayload.time = 1
	// 	} else if now.Hour() == 12 {
	// 		requestpayload.time = 2
	// 	} else if now.Hour() == 19 {
	// 		requestpayload.time = 3
	// 	} else {
	// 		continue
	// 	}
	err := ExecRequest(requestpayload)
	if err != nil {
		fmt.Println(err)
	}
	// }
}

// 	fmt.Println(now.String())
// }
