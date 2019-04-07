package main

import "memo_app/models"

var curriedId int

var memos models.MemoSlice

func init() {
	RepoCreateMemo(models.Memo{Title: "This is first memo", Content: "First content"})
	RepoCreateMemo(models.Memo{Title: "This is second memo", Content: "Second content"})
}

func RepoCreateMemo(m models.Memo) models.Memo {
	curriedId += 1
	m.Id = curriedId
	memos = append(memos, m)
	return m
}
