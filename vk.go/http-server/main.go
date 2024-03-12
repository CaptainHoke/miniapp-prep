package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

const keyServerAddress = "serverAddress"

func getRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Couldn't read body: %s\n", err)
	}

	fmt.Printf("%s: got / request. first(%t)=%s, second(%t)=%s, body:\n%s\n",
		ctx.Value(keyServerAddress),
		hasFirst, first,
		hasSecond, second, body)
	_, _ = io.WriteString(w, "Sex\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.PostFormValue("username")
	if username == "" {
		username = "Floppa"
	}

	fmt.Printf("%s: Got /hello request\n", ctx.Value(keyServerAddress))
	_, _ = io.WriteString(w, fmt.Sprintf("Ohayo, %s!\n", username))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx := context.Background()
	firstServer := &http.Server{Addr: ":3239", Handler: mux, BaseContext: func(l net.Listener) context.Context {
		ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
		return ctx
	}}

	err := firstServer.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("First server off")
	} else if err != nil {
		fmt.Printf("Error starting first server: %s\n", err)
		os.Exit(1)
	}
}
