package main

type CommandType string

const (
	ArithmeticCommand CommandType = "arithmetic"
	PushCommand       CommandType = "push"
	PopCommand        CommandType = "pop"
	LabelCommand      CommandType = "label"
	GotoCommand       CommandType = "goto"
	IfCommand         CommandType = "if"
	FunctionCommand   CommandType = "function"
	ReturnCommand     CommandType = "return"
	CallCommand       CommandType = "call"
)

var Segments = map[string]string{
	"argument": "ARG",
	"this":     "THIS",
	"that":     "THAT",
	"local":    "LCL",
	"temp":     "5",
	"static":   "16",
	"pointer":  "3",
}

var ArithCmds = map[string]string{
	"add": "M=D+M",
	"sub": "M=M-D",
	"neg": "M=-M",
	"and": "M=D&M",
	"or":  "M=D|M",
	"not": "M=!M",
}

var ComparisonCmds = map[string]string{
	"eq": "D;JEQ",
	"gt": "D;JGT",
	"lt": "D;JLT",
}
