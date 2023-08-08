import sys

from code_writer import CodeWriter
from enums import CommandType
from hack_parser import HackParser


class Translator:
    def __init__(self, file: str, filename: str) -> None:
        self.parser = HackParser(file=file)
        self.asm_file = open(f"{filename}.asm", "x")
        self.code_writer = CodeWriter(self.asm_file)

    def translate(self):
        while self.parser.has_more_lines:
            self.parser.advance()
            current_command = self.parser.get_current_command()

            if (
                current_command == CommandType.C_PUSH
                or current_command == CommandType.C_POP
            ):
                first_arg = self.parser.get_first_arg()
                second_arg = self.parser.get_second_arg()
                self.code_writer.write_push_pop_command(
                    command=current_command.value, segment=first_arg, index=second_arg
                )

            elif current_command == CommandType.C_ARITHMETIC:
                first_arg = self.parser.get_first_arg()
                self.code_writer.write_arithmetic_command(command=first_arg)
        self.code_writer.close()


if __name__ == "__main__":
    file = sys.argv[1]
    filename = file[: len(file) - 3]
    translator = Translator(file=file, filename=filename)
    translator.translate()
