package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Hello World! A message from kubernetes cloud."))
	})
	error := http.ListenAndServe(":8090", router)
	if error != nil {
		panic(error.Error())
	}
}
