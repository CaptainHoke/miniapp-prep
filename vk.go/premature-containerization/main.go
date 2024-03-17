package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func echo(args []string) error {
	if len(args) < 2 {
		return errors.New("no args to echo")
	}

	_, err := fmt.Println(strings.Join(args[1:], ", "))

	return err
}

func main() {
	if err := echo(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
