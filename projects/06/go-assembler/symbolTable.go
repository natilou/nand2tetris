package main

import "fmt"

type SymbolTable struct {
	tableMap	map[string]string
}

func NewSymbolTable() SymbolTable {
	symbols := make(map[string]string)
	for i := 0; i < 16; i++ {
		symbols[fmt.Sprintf("R%v", i)] = ConvertDecimalToBinary(i)
	}

	symbols["SP"] = ConvertDecimalToBinary(0)
	symbols["SP"] = ConvertDecimalToBinary(0)
	symbols["LCL"] = ConvertDecimalToBinary(1) 
	symbols["ARG"] = ConvertDecimalToBinary(2)
	symbols["THIS"] = ConvertDecimalToBinary(3)
	symbols["THAT"] = ConvertDecimalToBinary(4) 
	symbols["SCREEN"] = ConvertDecimalToBinary(16384)
	symbols["KBD"] = ConvertDecimalToBinary(24572)

	return SymbolTable{
		tableMap: symbols,
	}
}


func (s SymbolTable) AddEntry(symbol, address string) {
	s.tableMap[symbol] = address
}

func (s SymbolTable) Contains(symbol string) bool {
	_, isInTableMap := s.tableMap[symbol]
	return isInTableMap 
}

func (s SymbolTable) GetAddress(symbol string) string {
	return s.tableMap[symbol]
}