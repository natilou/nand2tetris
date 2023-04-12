
from utils import InstructionType


class HackParser:

    def __init__(self, file) -> None:
        self.filename = file.split('.')[0]
        self.read_file(file)
        self._instruction_type = ''
        self._dest = 'null'
        self._comp = '0'
        self._jump = 'null'
    
    @property
    def hack_file(self):
        self._hack_file = open(f'{self.filename}1.hack', 'x')
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
        self._instruction_type = ''

    @property
    def has_more_lines(self) -> bool:
        return self._asm_file_lines > self._lines_read

    def advance(self) -> None:
        current_line = self._asm_file.readline().strip()
        self._lines_read += 1

        if not current_line or current_line.startswith('/'):
            return

        elif current_line.startswith('@'):
            self._current_instruction = current_line
            self._instruction_type = InstructionType.A
            self.instruction_line +=1
      
        elif current_line.startswith('('):
            self._current_instruction = current_line
            self._instruction_type = InstructionType.L
        
        else:
            try:
                comment_index = current_line.index('/')
                self._current_instruction = current_line[:comment_index].strip()
            except:
                self._current_instruction = current_line

            self._instruction_type = InstructionType.C
            self._get_symbolic_c_instruction()
            self.instruction_line +=1
        
    def _get_symbolic_c_instruction(self) -> None:
        if ';' in self._current_instruction:
            self._comp, self._jump = self._current_instruction.split(';')
            self._dest = 'null'
        else:
            self._dest, self._comp = self._current_instruction.split('=')
            self._jump = 'null'

    def get_instruction_type(self) -> InstructionType:
        return self._instruction_type
    
    def get_symbol(self) -> str:
        if self._current_instruction.startswith('('):
            end_index = self._current_instruction.find(')')
            return self._current_instruction[1:end_index]
        else:
            return self._current_instruction[1:]
        
    def get_dest(self) -> str:
        return self._dest
    
    def get_comp(self) -> str:
        return self._comp
    
    def get_jump(self) -> str:
        return self._jump
        