package main

import (
	"log"

	"github.com/arekmano/deckard/cmd"
)

func main() {
	z := cmd.GetCommand()
	if err := z.Execute(); err != nil {
		log.Fatal(err)
	}
}
