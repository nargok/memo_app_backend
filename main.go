package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to memo App!")
}
