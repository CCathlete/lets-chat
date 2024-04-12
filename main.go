package main

import (
	"fmt"
	"log"
	"net/http"
)

type MessengerInput struct {
	Entry []struct {
		Time      uint `json:"time,omitempty"`
		Messaging []struct {
			Sender struct {
				Id string `json:"id"`
			} `json:"sender,omitempty"`
			Recipient struct {
				Id string `json:"id"`
			} `json:"recipient,omitempty"`
			Timestamp uint64 `json:"timestamp,omitempty"`
			Message   *struct {
				Mid  string `json:"mid,omitempty"`
				Seq  uint64 `json:"seq,omitempty"`
				Text string `json:"text"`
			} `json:"message,omitempty"`
		} `json:"messaging"`
	}
}

func main() {
	http.HandleFunc("/webhook", MessengerVerify)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func MessengerVerify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		challenge := r.URL.Query().Get("hub.challenge")
		verifyToken := r.URL.Query().Get("hub.verify_token")

		if len(verifyToken) > 0 && len(challenge) > 0 &&
			verifyToken == "developers-are-gods" {
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprint(w, challenge)
			return
		}

		w.Header().Set("content-Type", "text/plain")
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad Request")
	}
}
