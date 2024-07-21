package main

import (
	"log"
	"net/http"
	"time"

	"awesomeProject/cmd/rest/api"
	"awesomeProject/di"
)

// calls the rest api but using a local server instead of lambda

func main() {
	server := api.NewServer(di.New())

	r := http.NewServeMux()

	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler:     h,
		Addr:        "0.0.0.0:8080",
		ReadTimeout: 5 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
