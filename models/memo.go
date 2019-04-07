package models

import "time"

// +gen slice:"Where"
type Memo struct {
	Id 			int
	Title 		string
	Content 	string
	CreatedAt 	time.Time
}

type Memos []Memo