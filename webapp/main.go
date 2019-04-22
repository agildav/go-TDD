package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)

	const port = ":3000"

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("could not listen on port %s %v", port, err)
	}
}
