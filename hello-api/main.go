package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GreetingResponse struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var message string
	if name == "" {
		message = "Hello, World!"
	} else {
		message = fmt.Sprintf("Hello, %s!", name)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GreetingResponse{Message: message})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "OK"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /health", healthHandler)

	log.Printf("hello-api listening on :9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
