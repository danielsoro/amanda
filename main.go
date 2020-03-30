package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	phrases := [6]string{"Eu te amo!!",
		"Não vejo a hora de dormir contigo em meus braços todas as noites",
		"Já disse que tu é muito gostosa? Tu é MUITOO GOSTOSA!!",
		"Não existe princesa mais linda do que você, nem na Disney",
		"Minha tchutchuquinha, que eu amo muitcho!!!",
		"Já disse que tu é muito linda? Tu é MUITOO LINDA!!"}
	fmt.Fprint(w, phrases[rand.Intn(len(phrases))]);
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}