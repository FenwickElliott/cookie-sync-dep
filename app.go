package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8181"
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":"+port, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root got")
	io.WriteString(w, "Root directory")
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query())
}
