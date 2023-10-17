package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://localhost:8010/time")
		if err != nil {
			log.Fatalf("%v", err)
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Println(data)
		time.Sleep(3 * time.Second)
	}
}
