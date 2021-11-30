package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func calcHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	var operation operation
	err = json.Unmarshal(content, &operation)
	if err != nil {
		http.Error(w, "Error processing json", http.StatusInternalServerError)
		return
	}

	left := operation.Left
	right := operation.Right
	operator := operation.Operator

	fmt.Println("Requested to calculate", left, operator, right)

	var res int

	switch operator {
	case "+":
		res = Add(left, right)
	case "-":
		res = Subtract(left, right)
	case "*":
		res = Multiply(left, right)
	case "/":
		res = Divide(left, right)
	default:
		res = 0
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result{res})
}

type operation struct {
	Left     int    `json:"leftNumber"`
	Right    int    `json:"rightNumber"`
	Operator string `json:"operator"`
}

type result struct {
	Value int `json:"value"`
}
