package main

import (
    "fmt"
    "net/http"
)

func handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to the rewrite of the CitizenFleet Server using Go")
}

func main() {
	http.HandleFunc("/", handlerMain)
	http.ListenAndServe(":3000", nil)
}