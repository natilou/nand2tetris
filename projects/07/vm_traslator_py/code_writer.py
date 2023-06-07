from typing import Union

from enums import CommandType
from utils import get_comment, get_pop_cmd, get_push_cmd


class CodeWriter:
    """
    This module translates a parsed VM command into Hack assembly code.
    """

    def __init__(self, asm_file) -> None:
        self.asm_file = asm_file

    def write_arithmetic_command(self, command: str) -> None:
        self.asm_file.write(f"// ##  {command} \n")

    def write_push_pop_command(
        self,
        command: Union[CommandType.C_PUSH, CommandType.C_POP],
        segment: str,
        index: int,
    ) -> None:

        comment = get_comment(command, segment, index)
        asm_codes = get_pop_cmd(segment) if command == CommandType.C_POP else get_push_cmd(index)
        self.asm_file.write(f'{comment}\n')
        for code in asm_codes:
            self.asm_file.write(f'{code}\n')

    def close(self) -> None:
        pass
