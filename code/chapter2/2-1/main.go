package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.HandleFunc("/index", index)

	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world , i am index handler function")
}
