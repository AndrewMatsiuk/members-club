package main

import (
	"net/http"

	"members-club/handlers"
	"members-club/services"
)

func main() {
	members := services.New()
	handler := handlers.New(members)

	http.HandleFunc("/", handler.List)
	http.HandleFunc("/add", handler.Add)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
