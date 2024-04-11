package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/chatbot", myHandleFunc)
	http.ListenAndServe(":5000", nil)
}

func myHandleFunc(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// body is a []byte.
	userMessage := string(body)
	// Process userMessage and generate a response.
	response := processUserInput(userMessage)

	// Send the response back to the user.
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, response)
}

func processUserInput(input string) string {
	// I'm supposed to put chatbot logic here, wtf is it?
	return "Hello, I'm your chatbot."
}
