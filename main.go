package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
