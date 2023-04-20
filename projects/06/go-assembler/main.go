package main

import (
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	filename := argsWithoutProg[0]

	hackAssembler, err := NewHackAssembler(filename)
	
	if err != nil{
		log.Println("Error creating hack assembler:", err)
		return
	}

	hackAssembler.translate()
	
}
