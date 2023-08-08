package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type VMParser struct {
	currentCommand CommandType
	cachedLine     *string
	scanner        *bufio.Scanner
	currentLine    string
}

func NewVMParser(file string) (VMParser, error) {
	newFile, err := os.Open(file)
	if err != nil {
		fmt.Println("ERROR: Failed to create file:", err)
		return VMParser{}, err
	}

	return VMParser{
		currentCommand: "",
		cachedLine:     nil,
		scanner:        bufio.NewScanner(newFile),
		currentLine:    "",
	}, nil
}

func (p *VMParser) HasMoreLines() bool {
	if p.cachedLine != nil {
		return true
	}
	if !p.scanner.Scan() {
		return false
	}
	lineRead := p.scanner.Text()
	p.cachedLine = &lineRead
	return true
}

func (p *VMParser) Advance() {
	if p.scanner == nil {
		fmt.Println("Empty scanner, returning...")
		return
	}

	if p.cachedLine != nil {
		p.currentLine = *p.cachedLine
		p.cachedLine = nil
	} else {
		p.scanner.Scan()
		p.currentLine = p.scanner.Text()
	}

	p.currentLine = strings.TrimSpace(p.currentLine)
	operands := []string{"add", "sub", "neg", "gt", "eq", "lt", "or", "and", "not"}

	if p.currentLine == "" || strings.HasPrefix(p.currentLine, "/") {
		return
	} else if strings.Contains(p.currentLine, "push") {
		p.currentCommand = PushCommand
	} else if strings.Contains(p.currentLine, "pop") {
		p.currentCommand = PopCommand
	} else if strings.ContainsAny(p.currentLine, strings.Join(operands, "")) {
		p.currentCommand = ArithmeticCommand
	}
}

func (p *VMParser) GetFirstArg() string {
	if p.currentCommand == PushCommand || p.currentCommand == PopCommand || p.currentCommand == CallCommand || p.currentCommand == FunctionCommand {
		return strings.Split(p.currentLine, " ")[1]
	}
	return strings.Split(p.currentLine, " ")[0]
}

func (p *VMParser) GetSecondArg() string {
	return strings.Split(p.currentLine, " ")[2]
}

func (p *VMParser) GetCurrentCommand() CommandType {
	return p.currentCommand
}
