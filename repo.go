package main

var curriedId int

var memos Memos

func init() {
	RepoCreateMemo(Memo{Title: "This is first memo", Content: "First content"})
	RepoCreateMemo(Memo{Title: "This is second memo", Content: "Second content"})
}

func RepoCreateMemo(m Memo) Memo {
	curriedId += 1
	m.Id = curriedId
	memos = append(memos, m)
	return m
}

