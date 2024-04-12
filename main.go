package main

import (
	"log"
	"net/http"
)

const ()

func main() {
	http.HandleFunc("/webhook", MessengerVerify)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
