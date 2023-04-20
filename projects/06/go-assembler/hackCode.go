package main


type HackCode struct {
	destinationMap 	map[string]string
	computationMap 	map[string]string
	jumpMap 		map[string]string
}

func NewHackCode() HackCode{
	return HackCode{
		destinationMap: map[string]string{
			"null": "000",
			"M": "001",
			"D": "010",
			"DM": "011",
			"MD": "011",
			"A": "100",
			"AM": "101",
			"AD": "110",
			"ADM": "111",
		},
		computationMap: map[string]string{
			"0": "0101010",
			"1": "0111111",
			"-1": "0111010",
			"D": "0001100",
			"A": "0110000",
			"!D": "0001101",
			"!A": "0110001",
			"-D": "0001111",
			"-A": "0110011",
			"D+1": "0011111",
			"A+1": "0110111",
			"D-1": "0001110",
			"A-1": "0110010",
			"D+A": "0000010",
			"D-A": "0010011",
			"A-D": "0000111",
			"D&A": "0000000",
			"D|A": "0010101",
			"M": "1110000",
			"!M": "1110001",
			"-M": "1110011",
			"M+1": "1110111",
			"M-1": "1110010",
			"D+M": "1000010",
			"D-M": "1010011",
			"M-D": "1000111",
			"D&M": "1000000",
			"D|M": "1010101",
		},
		jumpMap: map[string]string{
			"null": "000",
			"JGT": "001",
			"JEQ": "010",
			"JGE": "011",
			"JLT": "100",
			"JNE": "101",
			"JLE": "110",
			"JMP": "111",
		
		},
	}
}

func (hC HackCode) GetBinaryDest(code string) string {
	return hC.destinationMap[code]
}

func (hC HackCode) GetBinaryComp(code string) string {
	return hC.computationMap[code]
}

func (hC HackCode) GetBinaryJump(code string) string {
	return hC.jumpMap[code]
}