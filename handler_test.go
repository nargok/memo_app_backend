package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v awnt %v", status, http.StatusOK)
	}

	expected := "Welcome to memo App!"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpedted body: got %v want %v", rr.Body.String(), expected)
	}


}
