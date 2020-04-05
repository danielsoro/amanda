package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Phrases struct {
	value string
}

func getPhrases() []Phrases {
	return []Phrases{
		Phrases{"Eu te amo!!"},
		Phrases{"Não vejo a hora de dormir contigo em meus braços todas as noites"},
		Phrases{"Já disse que tu é muito gostosa? Tu é MUITOO GOSTOSA!!"},
		Phrases{"Não existe princesa mais linda do que você, nem na Disney"},
		Phrases{"Minha tchutchuquinha, que eu amo muitcho!!!"},
		Phrases{"Já disse que tu é muito linda? Tu é MUITOO LINDA!!"},
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var phrases = getPhrases()
	fmt.Fprintf(w, "<h1>%s</h1>", phrases[rand.Intn(len(phrases))].value)
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
