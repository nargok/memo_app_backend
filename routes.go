package main

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MemoIndex",
		"GET",
		"/memos",
		MemoIndex,
	},
	Route{
		"MemoShow",
		"GET",
		"/memos/{memoId}",
		MemoRead,
	},
	Route{
		"MemoCreate",
		"POST",
		"/memos",
		MemoCreate,
		},
}
