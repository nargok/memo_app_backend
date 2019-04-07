package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wawnt %v", status, http.StatusOK)
	}

	expected := "Welcome to memo App!"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpedted body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestCreate(t *testing.T) {
	var memo = &Memo{
		Title: "テストタイトル",
		Content: "テストコンテンツ",
	}
	jsonMemo, _ := json.Marshal(memo)

	req, err := http.NewRequest("POST", "/memos", bytes.NewBuffer(jsonMemo))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MemoCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	lastMemo := memos[len(memos) - 1]
	if lastMemo.Title != memo.Title {
		t.Errorf("memo was not created since last memo title is %v want %v", lastMemo.Title ,memo.Title)
	}
}
