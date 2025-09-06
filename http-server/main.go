package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{store: NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5001", server))
}
