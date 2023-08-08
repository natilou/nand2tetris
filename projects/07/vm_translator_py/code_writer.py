from typing import Tuple

from constants import ARITHMETIC_CMDS, COMPARISON_CMDS, SEGMENTS
from enums import CommandType


class CodeWriter:
    """
    This module translates a parsed VM command into Hack assembly code.
    """

    def __init__(self, asm_file) -> None:
        self.asm_file = asm_file
        self.bool_count = 0

    def write_arithmetic_command(self, command: str) -> None:
        self._write_comment(command=command)

        if command != "not" and command != "neg":
            self._pop_stack_to_d()

        self._decrement_sp()
        self._set_a_register_to_stack()

        if command not in ["eq", "gt", "lt"]:
            self._write_arithmetic_cmd(command)

        elif command in ["eq", "gt", "lt"]:
            self._write(("D=M-D", f"@BOOL{self.bool_count}"))
            self._write_comparison_cmd(command)
            self._set_a_register_to_stack()
            self._write(("M=0", f"@ENDBOOL{self.bool_count}", "0;JMP"))
            self._write((f"(BOOL{self.bool_count})",))
            self._set_a_register_to_stack()
            self._write(("M=-1", f"(ENDBOOL{self.bool_count})"))
            self.bool_count += 1

        else:
            self._raise_unknown_command(command)

        self._increment_sp()

    def write_push_pop_command(
        self,
        command: str,
        segment: str,
        index: int,
    ) -> None:
        self._write_comment(command, segment, index)

        if command == CommandType.C_PUSH.value:
            self._push_cmd(segment, index)
            self._push_d_to_stack()

        elif command == CommandType.C_POP.value:
            self._pop_cmd(segment, index)

        else:
            self._raise_unknown_command(command)

    def _write_comment(self, command: str, segment: str = None, index: str = None):
        if not segment and not index:
            self.asm_file.write(f"// {command}\n")
        else:
            self.asm_file.write(f"// {command} {segment} {index}\n")

    def _write_arithmetic_cmd(self, command: str) -> None:
        self._write((ARITHMETIC_CMDS[command],))

    def _write_comparison_cmd(self, command: str) -> None:
        self._write((COMPARISON_CMDS[command],))

    def _write_segment_cmd(self, segment: str, index: str) -> None:
        asm_symbol = SEGMENTS.get(segment)
        if segment in ["pointer", "temp"]:
            self._write(
                (
                    f"@R{int(asm_symbol) + int(index)}",
                )
            )
        else:
            self._write(
                (
                    f"@{asm_symbol}",
                    "D=M",
                    f"@{index}",
                    "A=D+A",
                )
            )

    def _pop_cmd(self, segment: str, index: str):
        self._write_segment_cmd(segment, index)
        self._write(
            (
                "D=A",
                "@R13",
                "M=D",
            )
        )
        self._pop_stack_to_d()
        self._write(
            (
                "@R13",
                "A=M",
                "M=D",
            )
        )

    def _push_cmd(self, segment: str, index: str) -> None:
        if segment == "constant":
            self._write((f"@{index}", "D=A"))
        else:
            self._write_segment_cmd(segment, index)
            self._write(("D=M",))

    def _push_d_to_stack(self) -> None:
        asm_codes = (
            "@SP",
            "A=M",
            "M=D",
        )
        self._write(asm_codes=asm_codes)
        self._increment_sp()

    def _pop_stack_to_d(self) -> None:
        asm_codes = (
            "@SP",
            "M=M-1",
            "A=M",
            "D=M",
        )
        self._write(asm_codes=asm_codes)

    def _decrement_sp(self) -> None:
        asm_codes = ("@SP", "M=M-1")
        self._write(asm_codes=asm_codes)

    def _increment_sp(self) -> None:
        asm_codes = ("@SP", "M=M+1")
        self._write(asm_codes=asm_codes)

    def _set_a_register_to_stack(self) -> None:
        asm_codes = ("@SP", "A=M")
        self._write(asm_codes=asm_codes)

    def _write(self, asm_codes: Tuple[str]) -> None:
        for code in asm_codes:
            self.asm_file.write(f"{code}\n")

    def _raise_unknown_command(self, command: str):
        raise ValueError(f"{command} is an invalid command")

    def close(self) -> None:
        self.asm_file.close()
