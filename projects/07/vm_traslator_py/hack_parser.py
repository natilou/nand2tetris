from enums import CommandType


class HackParser:
    """
    This module handles the parsing of a single .vm file. The parser provides
    services for reading a VM command, unpacking the command into its
    various components, and providing convenient access to these components.
    In addition, the parser ignores all white space and comments.
    """

    def __init__(self, file: str) -> None:
        self.vm_file = open(file, "r")
        self.current_command: CommandType = None
        self.cached_line = None

    @property
    def has_more_lines(self) -> bool:
        if self.cached_line:
            return True

        line_read = self.vm_file.readline()
        if not line_read:
            return False

        self.cached_line = line_read.strip()
        return True

    def advance(self) -> None:
        if self.cached_line:
            self.raw_value = self.cached_line
            self.cached_line = None
        else:
            self.raw_value = self.vm_file.readline().strip()

        if not self.raw_value or self.raw_value.startswith("/"):
            return

        if "push" in self.raw_value:
            self.current_command = CommandType.C_PUSH

        elif "pop" in self.raw_value:
            self.current_command = CommandType.C_POP

        elif any(
            operand in self.raw_value
            for operand in ["add", "sub", "neg", "gt", "eq", "lt"]
        ):
            self.current_command = CommandType.C_ARITHMETIC

    def get_first_arg(self) -> str:
        if (
            self.current_command == CommandType.C_PUSH
            or self.current_command == CommandType.C_POP
            or self.current_command == CommandType.C_CALL
            or self.current_command == CommandType.C_FUNCTION
        ):
            return self.raw_value.split(" ")[1]

        return self.raw_value.split(" ")[0]

    def get_second_arg(self) -> str:
        return self.raw_value.split(" ")[2]

    def get_current_command(self) -> CommandType:
        return self.current_command
