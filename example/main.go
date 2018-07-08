package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/whaangbuu/hey"
)

func main() {
	r := hey.NewRouter()

	r.HandleFunc("GET", "/users/:id", userHandler)

	log.Fatalln(http.ListenAndServe(":3000", r))
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	userID := hey.Param(r.Context(), "id")

	fmt.Fprintf(w, userID)
}
