package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("please enter a command")
	}

	execute(args[1])
}
