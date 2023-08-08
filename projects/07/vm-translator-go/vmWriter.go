package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type VMWriter struct {
	asmFile   *os.File
	boolCount int
}

func NewVMWriter(asmFile *os.File) VMWriter {
	return VMWriter{
		asmFile:   asmFile,
		boolCount: 0,
	}
}

func (w *VMWriter) WriteArithmeticCommand(command string) {
	w.writeComment(command)
	if command != "not" && command != "neg" {
		w.popStackToD()
	}

	w.decrementSP()
	w.setARegisterToStack()

	if command != "eq" && command != "gt" && command != "lt" {
		w.writeArithmeticCommand(command)
	} else if command == "eq" || command == "gt" || command == "lt" {
		stringCount := strconv.Itoa(w.boolCount)
		w.write([]string{"D=M-D", fmt.Sprintf("@BOOL%s", stringCount)})
		w.writeComparisonCommand(command)
		w.setARegisterToStack()
		w.write([]string{"M=0", fmt.Sprintf("@ENDBOOL%s", stringCount), "0;JMP"})
		w.write([]string{fmt.Sprintf("(BOOL%s)", stringCount)})
		w.setARegisterToStack()
		w.write([]string{"M=-1", fmt.Sprintf("(ENDBOOL%s)", stringCount)})
		w.boolCount++
	} else {
		w.throwUnknownCommandError(command)
	}

	w.incrementSP()

}

func (w *VMWriter) WritePushPopCommand(command, segment, index string) {
	w.writeComment(command, map[string]string{"segment": segment, "index": index})

	if command == string(PushCommand) {
		w.pushCmd(segment, index)
		w.pushDToStack()
	} else if command == string(PopCommand) {
		w.popCmd(segment, index)
	} else {
		w.throwUnknownCommandError(command)
	}
}

func (w *VMWriter) writeComment(command string, params ...map[string]string) {

	if len(params) == 0 {
		w.asmFile.WriteString(fmt.Sprintf("// %s\n", command))
	} else {
		segment := ""
		index := ""

		paramsMap := params[0]
		if val, ok := paramsMap["segment"]; ok {
			segment = val
		}

		if val, ok := paramsMap["index"]; ok {
			index = val
		}

		w.asmFile.WriteString(fmt.Sprintf("// %s %s %s\n", command, segment, index))
	}
}

func (w *VMWriter) popStackToD() {
	asmCode := []string{
		"@SP",
		"M=M-1",
		"A=M",
		"D=M",
	}

	w.write(asmCode)
}

func (w *VMWriter) decrementSP() {
	asmCode := []string{"@SP", "M=M-1"}
	w.write(asmCode)
}

func (w *VMWriter) incrementSP() {
	asmCode := []string{"@SP", "M=M+1"}
	w.write(asmCode)
}

func (w *VMWriter) setARegisterToStack() {
	asmCode := []string{"@SP", "A=M"}
	w.write(asmCode)
}

func (w *VMWriter) writeArithmeticCommand(command string) {
	if asmCode, ok := ArithCmds[command]; ok {
		w.write([]string{asmCode})
	}
}

func (w *VMWriter) writeComparisonCommand(command string) {
	if asmCode, ok := ComparisonCmds[command]; ok {
		w.write([]string{asmCode})
	}
}

func (w *VMWriter) writeSegmentCmd(segment, index string) {
	if asmSymbol, ok := Segments[segment]; ok {
		if segment == "pointer" || segment == "temp" {
			address, err := strconv.Atoi(asmSymbol)
			if err != nil {
				log.Print("Can't convert string to int")
				panic(err)
			}
			indexValue, err := strconv.Atoi(index)
			if err != nil {
				log.Print("Can't convert string to int")
				panic(err)
			}
			value := address + indexValue
			w.write([]string{
				fmt.Sprintf("@R%d", value),
			})
		} else {
			w.write([]string{
				fmt.Sprintf("@%s", asmSymbol),
				"D=M",
				fmt.Sprintf("@%s", index),
				"A=D+A",
			})
		}
	}
}

func (w *VMWriter) pushCmd(segment, index string) {
	if segment == "constant" {
		w.write([]string{fmt.Sprintf("@%s", index), "D=A"})
	} else {
		w.writeSegmentCmd(segment, index)
		w.write([]string{"D=M"})
	}
}

func (w *VMWriter) popCmd(segment, index string) {
	w.writeSegmentCmd(segment, index)
	w.write([]string{
		"D=A",
		"@R13",
		"M=D",
	})
	w.popStackToD()
	w.write([]string{
		"@R13",
		"A=M",
		"M=D",
	})
}

func (w *VMWriter) pushDToStack() {
	asmCodes := []string{
		"@SP",
		"A=M",
		"M=D",
	}
	w.write(asmCodes)
	w.incrementSP()
}

func (w *VMWriter) write(asmCode []string) {
	for i := 0; i < len(asmCode); i++ {
		w.asmFile.WriteString(fmt.Sprintf("%s\n", asmCode[i]))
	}
}

func (w *VMWriter) throwUnknownCommandError(command string) error {
	return fmt.Errorf("%s is an invalid command", command)
}

func (w *VMWriter) Close() error {
	err := w.asmFile.Close()
	if err != nil {
		fmt.Printf("Error when closing the file: %s", err)
		return err
	}
	return nil
}
