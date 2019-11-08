package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	// func to return all articles written within memory
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("EndpointHit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// func to return string on the homepage
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequest() {
	// HandleFunc is a function out of http package that allows us to register functions with routes to handle what happens with that route.
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles) // INTERESTING THAT I DO NOT HAVE TO TELL THE HTTP ROUTER HANDLER WHICH KIND OF A REQUEST IS INCOMING
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
