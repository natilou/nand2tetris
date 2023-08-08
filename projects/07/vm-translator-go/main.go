package main

import (
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	filename := argsWithoutProg[0]

	vmTranslator, err := NewVMTranslator(filename)

	if err != nil {
		log.Println("Error creating vm translator:", err)
		return
	}

	vmTranslator.translate()

}
