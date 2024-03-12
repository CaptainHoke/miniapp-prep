package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / request")
	_, _ = io.WriteString(w, "Sex\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got /hello request")
	_, _ = io.WriteString(w, "I stole this code\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3239", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server off")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
