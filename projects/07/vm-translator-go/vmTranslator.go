package main

import (
	"fmt"
	"os"
)

type VMTranslator struct {
	parser VMParser
	writer VMWriter
}

func NewVMTranslator(file string) (VMTranslator, error) {
	parser, err := NewVMParser(file)
	if err != nil {
		fmt.Println("Error creating HackParser:", err)
		return VMTranslator{}, err
	}

	asmFile, err := os.Create(fmt.Sprintf("%v.asm", file[:len(file)-3]))
	if err != nil {
		fmt.Println("ERROR: Failed to create file:", err)
		return VMTranslator{}, err
	}

	writer := NewVMWriter(asmFile)

	return VMTranslator{
		parser: parser,
		writer: writer,
	}, nil
}

func (t *VMTranslator) translate() {
	for t.parser.HasMoreLines() {
		t.parser.Advance()
		currentCommand := t.parser.GetCurrentCommand()

		if currentCommand == PushCommand || currentCommand == PopCommand {
			firstArg := t.parser.GetFirstArg()
			secondArg := t.parser.GetSecondArg()
			t.writer.WritePushPopCommand(string(currentCommand), firstArg, secondArg)
		} else if currentCommand == ArithmeticCommand {
			firstArg := t.parser.GetFirstArg()
			t.writer.WriteArithmeticCommand(firstArg)

		}
	}
	t.writer.Close()
}
