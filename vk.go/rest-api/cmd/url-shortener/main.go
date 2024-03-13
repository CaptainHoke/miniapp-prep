package main

import (
	"fmt"
	"rest-api/core/config"
)

func main() {

	cfg := config.MustLoadCfg()

	fmt.Println(cfg)

	// TODO: logger (slog)
	// TODO: db (postgres)
	// TODO: router (go-chi for now => custom impl later)
	// TODO: run the server

}
