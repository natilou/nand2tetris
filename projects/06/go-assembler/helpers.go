package main

import (
	"fmt"
)

type InstructionType string

const (
	InstructionA InstructionType = "INSTRUCTION TYPE A"
	InstructionC InstructionType = "INSTRUCTION TYPE C"
	InstructionL InstructionType = "INSTRUCTION TYPE L"
)

func ConvertDecimalToBinary(decimal int) string {
	return fmt.Sprintf("%016b", decimal)
}
