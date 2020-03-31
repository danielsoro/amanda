package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	phrases := [6]string{"Eu te amo!!",
		"Não vejo a hora de dormir contigo em meus braços todas as noites",
		"Já disse que tu é muito gostosa? Tu é MUITOO GOSTOSA!!",
		"Não existe princesa mais linda do que você, nem na Disney",
		"Minha tchutchuquinha, que eu amo muitcho!!!",
		"Já disse que tu é muito linda? Tu é MUITOO LINDA!!"}
	fmt.Fprintf(w, "<h1>%s</h1>", phrases[rand.Intn(len(phrases))])
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Print("$PORT == 8080")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
