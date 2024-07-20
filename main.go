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

	// Write the response
	fmt.Fprint(w, message)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/frases", quotesHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
