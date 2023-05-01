from instruction import Instruction
from utils import InstructionType


class HackParser:

    def __init__(self, file) -> None:
        self.filename = file.split('.')[0]
        self.read_file(file)
        self.current_instruction = Instruction()

    @property
    def hack_file(self):
        self._hack_file = open(f'{self.filename}-py.hack', 'x')
        return self._hack_file

    def read_file(self, file) -> None:
        self._asm_file = open(file, "r")
        self._asm_file_lines = len(self._asm_file.readlines())
        self._asm_file.seek(0)
        self._lines_read = 0
        self.instruction_line = -1
    
    def reset_file(self):
        self._asm_file.seek(0)
        self._lines_read = 0
        self.current_instruction = Instruction()

    @property
    def has_more_lines(self) -> bool:
        return self._asm_file_lines > self._lines_read

    def advance(self) -> None:
        current_line = self._asm_file.readline().strip()
        self._lines_read += 1

        self.current_instruction.process_instruction(current_line)
        instruction_type = self.get_instruction_type()
        if instruction_type == InstructionType.A or instruction_type == InstructionType.C:
            self.instruction_line += 1
        
    def get_instruction_type(self) -> InstructionType:
        return self.current_instruction.get_instruction_type()
    
    def get_symbol(self) -> str:
        return self.current_instruction.get_symbol()
        
    def get_dest(self) -> str:
        return self.current_instruction.get_destination()
    
    def get_comp(self) -> str:
        return self.current_instruction.get_computation()
    
    def get_jump(self) -> str:
        return self.current_instruction.get_jump()
       