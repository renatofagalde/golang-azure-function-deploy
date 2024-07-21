package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

var quotes = []string{
	"Eu não tenho ídolos. Tenho admiração por trabalho, dedicação e competência.",
	"O importante é ganhar. Tudo e sempre. Essa história que o importante é competir não passa de demagogia.",
	"Sem sacrifício, não há vitória.",
	"Seu amor me fortalece, seu ódio me motiva.",
}

func quotesHandler(w http.ResponseWriter, r *http.Request) {
	// Get a random quote
	message := quotes[rand.Intn(len(quotes))]

	response := fmt.Sprintf("-> %s <-", message)

	// Write the response
	fmt.Fprint(w, response)
	log.Printf("Response sent: %s", response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	fmt.Println("azure function 02")
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
		log.Printf("Port from environment variable: %s", listenAddr)
	} else {
		log.Println("Using default port: 8080")
	}
	http.HandleFunc("/api/frases", quotesHandler)
	http.HandleFunc("/health", healthHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s/", listenAddr, listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
