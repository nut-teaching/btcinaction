package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello world")
	})

	http.HandleFunc("/index", indexNew)

	http.ListenAndServe(":8090", nil)
}

func indexNew(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "i am index function")
}
