import sys
from hack_code import HackCode
from hack_parser import HackParser
from utils import (
    convert_decimal_to_binary,
    InstructionType
)


class HackAssembler:
    
    def __init__(self, file: str) -> None:
        self.parser: HackParser = HackParser(file=file)
        self.hack_file = self.parser.hack_file
    
    def translate(self) -> None:

        while self.parser.has_more_lines:
            self.parser.advance()
            current_instruction = self.parser.get_instruction_type()

            if current_instruction:
                
                if current_instruction == InstructionType.A:
                    symbol = self.parser.get_symbol()
                    binary = convert_decimal_to_binary(int(symbol))
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



if __name__ == "__main__":
    file = sys.argv[1]
    hack_assembler = HackAssembler(file=file)
    hack_assembler.translate()
