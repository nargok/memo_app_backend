package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"memo_app/models"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to memo App!")
}

func MemoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(memos); err != nil {
		panic(err)
	}
}

func MemoRead(w http.ResponseWriter, r *http.Request) {
	// request内容からidを取り出す
	vars := mux.Vars(r)
	memoId, _ := strconv.Atoi(vars["memoId"])

	targetMemo := memos.Where(func(memo models.Memo) bool {
		return memo.Id == memoId
	})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if targetMemo == nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(targetMemo); err != nil {
		panic(err)
	}
}

func MemoCreate(w http.ResponseWriter, r *http.Request) {
	var memo models.Memo
	memo.CreatedAt = time.Now()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &memo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	m := RepoCreateMemo(memo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}