package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().Unix())

	// Define your HTTP handler function
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		// Choose a random message from the slice
		text := messages[rand.Intn(len(messages))]

		// Create a Message object
		message := Message{Text: text}

		// Marshal the message into JSON
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the content type header and write the JSON response
		w.Header().Set("Content-Type", "application/json")

		// Set the CORS headers to allow cross-origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.Write(jsonMessage)
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
