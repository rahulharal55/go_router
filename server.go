package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func createServer() {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Printf("All Query Params : %+v", q)
	fmt.Println(q.Get("a"))
	fmt.Println(q.Get("b"))
	fmt.Println(q.Get(""))
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
