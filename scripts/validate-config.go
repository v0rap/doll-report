package main

import (
	"os"

	"github.com/TwiN/gatus/v5/config"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		os.Exit(2)
	}
	_, err := config.LoadConfiguration(args[1])
	if err != nil {
		os.Exit(3)
	}
}
