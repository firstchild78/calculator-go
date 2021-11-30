package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendRequest(op operation) (int, int) {

	url := "http://localhost:1234/calculate"

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(op)

	req, err := http.Post(url, "application/json", body)
	if err != nil {
		fmt.Println("Could not Get", err)
		return 0, 503
	}

	res := result{}
	json.NewDecoder(req.Body).Decode(&res)
	if err != nil {
		fmt.Println("Could not parse request", err)
		return 0, 500
	}

	return res.Value, req.StatusCode
}

type operation struct {
	Left     int    `json:"leftNumber"`
	Right    int    `json:"rightNumber"`
	Operator string `json:"operator"`
}

type result struct {
	Value int `json:"value"`
}
