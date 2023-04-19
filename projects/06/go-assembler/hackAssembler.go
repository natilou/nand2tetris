package main

import (
	"fmt"
	"os"
	"strconv"
)

type HackAssembler struct {
	parser      		HackParser
	hackFile    		*os.File
	symbolTable 		SymbolTable
	hackCode    		HackCode
	symbolRamAddress 	int
}

func NewHackAssembler(file string) (HackAssembler, error) {
	parser, err := NewHackParser(file)
	if err != nil {
		fmt.Println("Error creating HackParser:", err)
		return HackAssembler{}, err
	}

	return HackAssembler{
		parser:      parser,
		hackFile:    parser.HackFile,
		symbolTable: NewSymbolTable(),
		hackCode:    NewHackCode(),
		symbolRamAddress: 16,
	}, nil
}

func (ha *HackAssembler) translate() {
	// first read
	ha.parser.ReadFile()
	for ha.parser.HasMoreLines() {
		ha.parser.Advance()
		currentInstruction := ha.parser.GetInstructionType()

		if currentInstruction == InstructionL {
			symbol := ha.parser.GetSymbol()
			address := ConvertDecimalToBinary(ha.parser.InstructionLine + 1)
			ha.symbolTable.AddEntry(symbol, address)
		}

	}

	// second read
	ha.parser.ResetFile()
	for ha.parser.HasMoreLines() {

		ha.parser.Advance()
		currentInstruction := ha.parser.GetInstructionType()

		if currentInstruction == InstructionA {
			symbol := ha.parser.GetSymbol()
			binary := ha.processAInstruction(symbol)
			ha.hackFile.WriteString(fmt.Sprintf("%s\n", binary))

		} else if currentInstruction == InstructionC {
			destination := ha.parser.GetDest()
			computation := ha.parser.GetComp()
			jump := ha.parser.GetJump()

			binary_dest := ha.hackCode.GetBinaryDest(destination)
			binary_comp := ha.hackCode.GetBinaryComp(computation)
			binary_jump := ha.hackCode.GetBinaryJump(jump)

			ha.hackFile.WriteString(fmt.Sprintf("111%s%s%s\n", binary_comp, binary_dest, binary_jump))
		}
	}
}

func (ha *HackAssembler) processAInstruction(symbol string) string {
	var binary string

	if symbolValue, err := strconv.Atoi(symbol); err == nil {
		binary = ConvertDecimalToBinary(symbolValue)
	} else {
		if !ha.symbolTable.Contains(symbol) {
			binaryAddress := ConvertDecimalToBinary(ha.symbolRamAddress)
			ha.symbolTable.AddEntry(symbol, binaryAddress)
			ha.symbolRamAddress++
		}
		binary = ha.symbolTable.GetAddress(symbol)
	}
	return binary
}
