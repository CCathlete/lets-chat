package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const PAGE_TOKEN = ""

type MessengerInput struct {
	Entry []struct {
		Time     uint `json:"time,omitempty"`
		Messages []struct {
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
	} else if r.Method == "POST" {
		defer r.Body.Close()

		input := new(MessengerInput)
		if err := json.NewDecoder(r.Body).Decode(input); err == nil {
			log.Println("got message:",
				input.Entry[0].Messages[0].Message.Text)
			/* The clients POST request is the message they send.
			We would like our letschat bot to reply. */
			reply := input.Entry[0].Messages[0]
			/* We're now replying so we need to swap the sender
			and reciever. */
			reply.Sender, reply.Recipient = reply.Recipient, reply.Sender

			if replyText, err := getApiAiResponse(*input); err == nil {
				reply.Message.Text = replyText
				// reply.Message.Seq // What is the purpose of this?!
				// reply.Message.Mid
				replyBody, _ := json.Marshal(reply)
				url := fmt.Sprintf("https://graph.facebook.com/v2.6/me/"+
					"messages?access_token=%s", PAGE_TOKEN)

				http.Post(url, "application/json", bytes.NewReader(replyBody))
			}
			return
		} else {
			log.Fatalln("Couldn't decode message's body: ", err)
		}
	}
	w.Header().Set("content-Type", "text/plain")
	w.WriteHeader(400)
	fmt.Fprint(w, "Bad Request")
}
