package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type HackParser struct {
	filename           string
	HackFile           *os.File
	asmFile            *os.File
	cachedLine         *string
	InstructionLine    int
	scanner            *bufio.Scanner
	currentInstruction Instruction
}

func NewHackParser(file string) (HackParser, error) {
	newFile, err := os.Create(fmt.Sprintf("%v-go.hack", file[:len(file)-4]))
	if err != nil {
		fmt.Println("ERROR: Failed to create file:", err)
		return HackParser{}, err
	}

	return HackParser{
		filename:           file,
		HackFile:           newFile,
		InstructionLine:    -1,
		currentInstruction: NewEmptyInstruction(),
	}, nil
}

func (h *HackParser) ReadFile() error {
	file, err := os.Open(h.filename)
	if err != nil {
		fmt.Println("ERROR: Failed to open file:", err)
		return err
	}

	h.asmFile = file
	h.scanner = bufio.NewScanner(h.asmFile)
	return nil
}

func (h *HackParser) ResetFile() {
	h.asmFile.Seek(0, 0)
	h.scanner = bufio.NewScanner(h.asmFile)
	h.cachedLine = nil
	h.currentInstruction = NewEmptyInstruction()
}

func (h *HackParser) HasMoreLines() bool {
	if h.cachedLine != nil {
		return true
	}
	if !h.scanner.Scan() {
		return false
	}
	currentLine := h.scanner.Text()
	h.cachedLine = &currentLine
	return true
}

func (h *HackParser) Advance() {
	if h.scanner == nil {
		fmt.Println("Empty scanner, returning...")
		return
	}

	var currentLine string
	if h.cachedLine != nil {
		currentLine = *h.cachedLine
		h.cachedLine = nil
	} else {
		h.scanner.Scan()
		currentLine = h.scanner.Text()
	}

	currentLine = strings.TrimSpace(currentLine)
	h.currentInstruction.ProcessInstruction(currentLine)
	if h.GetInstructionType() == InstructionA || h.GetInstructionType() == InstructionC {
		h.InstructionLine++
	}
}

func (h *HackParser) GetInstructionType() InstructionType {
	return h.currentInstruction.GetInstructionType()
}

func (h *HackParser) GetSymbol() string {
	return h.currentInstruction.GetSymbol()
}

func (h *HackParser) GetDest() string {
	return h.currentInstruction.GetDestination()
}

func (h *HackParser) GetComp() string {
	return h.currentInstruction.GetComputation()
}

func (h *HackParser) GetJump() string {
	return h.currentInstruction.GetJump()
}
