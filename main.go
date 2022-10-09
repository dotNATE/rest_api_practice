package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Projects []Project

func allProjects(w http.ResponseWriter, r *http.Request) {
	projects := Projects{
		Project{Title: "Example project", Description: "This is a description", Link: "https://github.com/dotNATE/rest_api_practice"},
	}

	fmt.Println("Endpoint hit: All projects endpoint")
	json.NewEncoder(w).Encode(projects)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/projects", allProjects)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	handleRequests()
}
