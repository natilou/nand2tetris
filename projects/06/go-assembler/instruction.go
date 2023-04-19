package main

import (
	"strings"
)

type Instruction struct {
	instructionType InstructionType
	destination     string
	computation     string
	jump            string
	rawValue        string
}

func NewEmptyInstruction() Instruction {
	return Instruction{}
}

func (i *Instruction) ProcessInstruction(rawValue string) {
	i.rawValue = rawValue

	if i.rawValue == "" || strings.HasPrefix(i.rawValue, "/") {
		return
	} else if strings.HasPrefix(i.rawValue, "@") {
		i.instructionType = InstructionA
	} else if strings.HasPrefix(i.rawValue, "(") {
		i.instructionType = InstructionL
	} else {
		i.instructionType = InstructionC
		i.rawValue = strings.Split(i.rawValue, "//")[0]
		i.getSymbolicCInstruction()
	}
}

func (i *Instruction) getSymbolicCInstruction() {
	if strings.Contains(i.rawValue, ";") {
		parts := strings.Split(i.rawValue, ";")
		i.computation = strings.TrimSpace(parts[0])
		i.jump = strings.TrimSpace(parts[1])
		i.destination = "null"
	} else {
		parts := strings.Split(i.rawValue, "=")
		i.computation = strings.TrimSpace(parts[1])
		i.destination = strings.TrimSpace(parts[0])
		i.jump = "null"
	}
}

func (i *Instruction) GetSymbol() string {
	if strings.HasPrefix(i.rawValue, "(") {
		endIndex := strings.Index(i.rawValue, ")")
		return i.rawValue[1:endIndex]
	}
	return i.rawValue[1:]
}

func (i *Instruction) GetDestination() string {
	return i.destination
}

func (i *Instruction) GetComputation() string {
	return i.computation
}

func (i *Instruction) GetJump() string {
	return i.jump
}

func (i *Instruction) GetInstructionType() InstructionType {
	return i.instructionType
}
