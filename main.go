package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	handleRequests()
}
