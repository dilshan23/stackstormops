package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Type      string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

type ResponseBody struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Body  struct {
		Challenge string `json:"challenge"`
	} `json:"body"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Process the request
	if reqBody.Type == "url_verification" {
		// Prepare the response
		resBody := ResponseBody{
			Code:  http.StatusOK,
			Error: "challenge_failed",
			Body: struct {
				Challenge string `json:"challenge"`
			}{
				Challenge: reqBody.Challenge,
			},
		}

		// Set the response headers
		w.Header().Set("Content-Type", "application/json")

		// Send the response
		err = json.NewEncoder(w).Encode(resBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		// Invalid request type
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

