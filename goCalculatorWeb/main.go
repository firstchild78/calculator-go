package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":1235", nil)
	if err != nil {
		log.Fatal("Fail to listen on port", err.Error())
	}
}
