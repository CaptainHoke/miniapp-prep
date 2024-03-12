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

	fmt.Printf("%s: Got / request\n", ctx.Value(keyServerAddress))
	_, _ = io.WriteString(w, "Sex\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: Got /hello request\n", ctx.Value(keyServerAddress))
	_, _ = io.WriteString(w, "I stole this code\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())
	firstServer := &http.Server{Addr: ":3239", Handler: mux, BaseContext: func(l net.Listener) context.Context {
		ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
		return ctx
	}}
	secondServer := &http.Server{Addr: ":6969", Handler: mux, BaseContext: func(l net.Listener) context.Context {
		ctx = context.WithValue(ctx, keyServerAddress, l.Addr().String())
		return ctx
	}}

	go func() {
		err := firstServer.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("First server off")
		} else if err != nil {
			fmt.Printf("Error starting first server: %s\n", err)
			os.Exit(1)
		}

		cancelCtx()
	}()

	go func() {
		err := secondServer.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("Second server off")
		} else if err != nil {
			fmt.Printf("Error starting second server: %s\n", err)
			os.Exit(1)
		}

		cancelCtx()
	}()

	<-ctx.Done()
}
