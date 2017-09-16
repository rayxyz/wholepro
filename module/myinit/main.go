package main

import "net/http"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, there!"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":9999", nil)
}
