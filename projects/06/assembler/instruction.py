from utils import InstructionType


class Instruction:

    def __init__(self) -> None:
        self.instruction_type: InstructionType = ''
        self.destination: str = 'null'
        self.computation: str = '0'
        self.jump: str = 'null'
        self.raw_value: str = ''

    def process_instruction(self, raw_value: str) -> None:
        self.raw_value = raw_value
        
        if not self.raw_value or self.raw_value.startswith('/'):
            return

        elif self.raw_value.startswith('@'):
            self.instruction_type = InstructionType.A
      
        elif self.raw_value.startswith('('):
            self.instruction_type = InstructionType.L
        
        else:
            self.instruction_type = InstructionType.C
            self.raw_value = self.raw_value.split("//")[0]
            self._get_symbolic_c_instruction()
    
    def _get_symbolic_c_instruction(self) -> None:

        if ';' in self.raw_value:
            parts = self.raw_value.split(';')
            self.computation = parts[0].strip()
            self.jump = parts[1].strip()
            self.destination = 'null'
        else:
            parts = self.raw_value.split('=')
            self.computation = parts[1].strip()
            self.destination = parts[0].strip()
            self.jump = 'null'
        
    def get_symbol(self) -> str:
        if self.raw_value.startswith('('):
            end_index = self.raw_value.find(')')
            return self.raw_value[1:end_index]
     
        return self.raw_value[1:]

    def get_destination(self) -> str:
        return self.destination

    def get_computation(self) -> str:
        return self.computation
    
    def get_jump(self) -> str:
        return self.jump

    def get_instruction_type(self) -> InstructionType:
        return self.instruction_type
