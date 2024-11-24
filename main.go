package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string

type requestBody struct {
	Message string `json:"message"`
}

// Handler for the GET request
func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", message)
}

// Handler for the POST request
// {"message": "task1"}
func postHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	message = reqBody.Message
	fmt.Fprintln(w, "Message updated successfully!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/get", getHandler).Methods("GET")
	router.HandleFunc("/api/post", postHandler).Methods("POST")

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
