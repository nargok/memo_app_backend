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
		"TodoIndex",
		"GET",
		"/memos",
		MemoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/memos",
		MemoCreate,
		},
}
