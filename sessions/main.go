package main

import (
	"encoding/gob"
	"net/http"
	"test-modules/sessions/handlers"
	"test-modules/sessions/model"

	"github.com/gorilla/mux"
)

func main() {
	gob.Register(model.SessionData{})

	router := mux.NewRouter()
	router.StrictSlash(true).Path("/").HandlerFunc(handlers.IndexPageHandler()).Methods("GET")

	router.StrictSlash(true).Path("/name").HandlerFunc(handlers.SetNameHandler()).Methods("POST")
	router.StrictSlash(true).Path("/name").HandlerFunc(handlers.GetNameHandler()).Methods("GET")

	router.StrictSlash(true).Path("/number").HandlerFunc(handlers.SetNumberHandler()).Methods("POST")
	router.StrictSlash(true).Path("/number").HandlerFunc(handlers.GetNumberHandler()).Methods("GET")

	http.ListenAndServe(":80", router)
}
