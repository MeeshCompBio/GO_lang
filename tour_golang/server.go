package main

import (
	"fmt"
	"net/http"
)

// simple function to gather response
func helloSWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello sever person")
}

// handle and read
func server() {
	http.HandleFunc("/", helloSWorld)
	http.ListenAndServe(":8123", nil)
}
