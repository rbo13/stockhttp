package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/whaangbuu/hey"
)

// Resp is a struct that represents
// a response.
type Resp struct {
	Success    bool        `json:"success,omitempty"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

func main() {
	r := hey.NewRouter()

	r.HandleFunc("GET", "/users/:id", userHandler)
	r.HandleFunc("POST", "/user/create", userCreateHandler)
	log.Fatalln(http.ListenAndServe(":3000", r))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := hey.Param(r.Context(), "id")

	fmt.Fprintf(w, userID)
}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	// From an html form data
	err := r.ParseForm()

	if err != nil {
		log.Printf("Error parsing form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := Resp{
		Success:    true,
		StatusCode: http.StatusCreated,
		Data:       username,
	}
	json.NewEncoder(w).Encode(resp)
}
