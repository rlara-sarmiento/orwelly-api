package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rlara-sarmiento/orwelly-api/pkg/api/router"
)

func main() {
	log.Println("Starting orwelly-api")

	address := "localhost:8080"
	wt := 15 * time.Second
	rt := 15 * time.Second

	server := &http.Server{
		Handler:      router.Get(),
		Addr:         address,
		WriteTimeout: wt,
		ReadTimeout:  rt,
	}

	log.Printf("Listening at http://%s\n", address)
	log.Fatal(server.ListenAndServe())
}
