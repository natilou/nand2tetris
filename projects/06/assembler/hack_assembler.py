import sys
from hack_code import HackCode
from hack_parser import HackParser
from symbol_table import SymbolTable
from utils import (
    convert_decimal_to_binary,
    InstructionType,
)


class HackAssembler:
    
    def __init__(self, file: str) -> None:
        self.parser: HackParser = HackParser(file=file)
        self.hack_file = self.parser.hack_file
        self.symbol_table: SymbolTable = SymbolTable()
        self.symbol_ram_address = 16
    
    def translate(self) -> None:

        # first pass
        while self.parser.has_more_lines:
            self.parser.advance()
            current_instruction = self.parser.get_instruction_type()

            if current_instruction == InstructionType.L:
                symbol = self.parser.get_symbol()
                address = convert_decimal_to_binary(self.parser.instruction_line+1)
                self.symbol_table.add_entry(symbol, address)
        
        # second pass
        self.parser.reset_file()
        while self.parser.has_more_lines:

            self.parser.advance()
            current_instruction = self.parser.get_instruction_type()
             
            if current_instruction == InstructionType.A:
                symbol = self.parser.get_symbol()
                binary = self.process_a_instruction(symbol)
                self.hack_file.write(f'{binary}\n')

            elif current_instruction == InstructionType.C:
                dest = self.parser.get_dest()
                comp = self.parser.get_comp()
                jump = self.parser.get_jump()

                binary_dest = HackCode.get_binary_dest(dest) 
                binary_comp = HackCode.get_binary_comp(comp) 
                binary_jump = HackCode.get_binary_jump(jump) 

                self.hack_file.write(f'111{binary_comp}{binary_dest}{binary_jump}\n')
            
        self.hack_file.close()

    def process_a_instruction(self, symbol: str) -> str:
        if symbol.isdigit():
            binary = convert_decimal_to_binary(int(symbol))
        else:
            if not self.symbol_table.contains(symbol):
                binary_address = convert_decimal_to_binary(int(self.symbol_ram_address))
                self.symbol_table.add_entry(symbol, binary_address)
                self.symbol_ram_address += 1
            binary = self.symbol_table.get_address(symbol)
        return binary



if __name__ == "__main__":
    file = sys.argv[1]
    hack_assembler = HackAssembler(file=file)
    hack_assembler.translate()
