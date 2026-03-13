package main

import (
	"os"
	"log"
	"io"
	"github.com/TwiN/gatus/v5/config"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	bytes, _ := io.ReadAll(os.Stdin)
	// log.Println(err, string(bytes))

	f, err := os.CreateTemp("", "temp.yaml")
	check(err)

	defer os.Remove(f.Name())

	_, err = f.Write(bytes)
	check(err)

	/*
	args := os.Args
	if len(args) != 2 {
		os.Exit(2)
	}
	*/
	// _, err := config.LoadConfiguration(args[1])
	_, err = config.LoadConfiguration(f.Name())
	if err != nil {
		log.Fatal(err);
		os.Exit(3)
	}
}
