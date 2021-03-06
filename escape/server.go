package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, this is home page")
}

func getUserName(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("name")
	log.Printf("name = %s", v)
	fmt.Fprintf(w, v)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/getUserName", getUserName)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
