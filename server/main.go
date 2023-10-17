package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", func(w http.ResponseWriter, req *http.Request) { 
		fmt.Fprintln(w, time.Now().String())})
	if err := http.ListenAndServe(":8010", nil); err != nil {
		log.Fatalf("%v", err)
	}
}
