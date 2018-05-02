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
		port = "80"
	}

	fmt.Println("Setting up, listening on (internal) port:", port)

	http.HandleFunc("/", root)
	http.ListenAndServe(":"+port, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("root got")
	io.WriteString(w, "Root")
}
