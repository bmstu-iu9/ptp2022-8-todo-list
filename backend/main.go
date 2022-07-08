package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

func main() {
	http.HandleFunc("/hello", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server start failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	counter++
	_, err := w.Write([]byte(fmt.Sprintf("Георгий, вот ваш API. (%v)", counter)))
	if err != nil {
		log.Fatal("Failed to write response")
	}
}
