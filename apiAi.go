package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const AUTH_TOKEN = ""

type ApiAiInput struct {
	Status struct {
		Code      int
		ErrorType string
	}
	Result struct {
		Action           *string
		ActionIncomplete bool
		Speech           string
	} `json:"result"`
}

func getApiAiResponse(msgInpt MessengerInput) (
	replyText string, err error,
) {
	params := url.Values{}
	params.Add("query", msgInpt.Entry[0].Messages[0].Message.Text)
	params.Set("sessionId", msgInpt.Entry[0].Messages[0].Sender.Id)

	url := fmt.Sprintf("https://api.api.ai/v1/query?V=20160518&lang=En%s",
		params.Encode())
	ai, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	ai.Header.Set("Authorization", "Bearer "+AUTH_TOKEN)

	if resp, err := http.DefaultClient.Do(ai); err != nil {
		return "", err
	} else {
		defer resp.Body.Close()

		var aiInput ApiAiInput
		datastring, _ := io.ReadAll(resp.Body)
		err := json.NewDecoder(strings.NewReader(
			string(datastring))).Decode(&aiInput)
		if err != nil {
			return "", err
		}
		return aiInput.Result.Speech, nil
	}
}
