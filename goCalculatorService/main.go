package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Calculator Service Starting...")
	err := http.ListenAndServe(":1234", handler())
	if err != nil {
		log.Fatal("Fail to listen on port", err.Error())
	}
}

func handler() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/calculate", calcHandler)
	return r
}
