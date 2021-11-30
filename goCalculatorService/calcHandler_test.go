package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegBlackBox(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	v := []struct {
		left      int
		right     int
		operation string
		result    int
	}{
		{2, 4, "+", 6},
		{2, 4, "-", -2},
		{2, 4, "*", 8},
		{4, 2, "/", 2},
		{4, 0, "/", 0},
	}
	for _, v := range v {
		op := operation{v.left, v.right, v.operation}
		body := new(bytes.Buffer)
		json.NewEncoder(body).Encode(op)

		req, err := http.Post(fmt.Sprintf("%s/calculate", srv.URL), "application/json", body)
		if err != nil {
			t.Error("Could not Get", err)
		}

		assert.Equal(t, http.StatusOK, req.StatusCode)

		res := result{}
		json.NewDecoder(req.Body).Decode(&res)
		if err != nil {
			t.Fatal("Could not parse request", err)
		}
		assert.Equal(t, v.result, res.Value)
	}
}

func TestIntegBadData(t *testing.T) {
	srv := httptest.NewServer(handler())
	defer srv.Close()

	body := bytes.NewReader([]byte("{\"bad-json\": 5"))
	req, err := http.Post(fmt.Sprintf("%s/calculate", srv.URL), "application/json", body)
	if err != nil {
		t.Error("Could not Get", err)
	}
	assert.Equal(t, http.StatusInternalServerError, req.StatusCode)
}
