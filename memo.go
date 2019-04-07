package main

import "time"

type Memo struct {
	Id 			int
	Title 		string
	Content 	string
	CreatedAt 	time.Time
}

type Memos []Memo