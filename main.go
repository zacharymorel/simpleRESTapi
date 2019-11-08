package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST POST ENDPOINT WORKED")
}

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
	myRouter := mux.NewRouter().StrictSlash(true)
	// WE USE gorilla/mux router so we can specify what verbs can be used for what routes with a much simpler syntax

	// HandleFunc is a function out of http package that allows us to register functions with routes to handle what happens with that route.
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
