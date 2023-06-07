RAM_ADDRESS = {
    "SP": "0",
    "local": "1",
    "argument": "2",
    "this": "3",  # pointer 0
    "that": "4",  # pointer 1
    "temp": "5",  # from 5 - 12
    "R13": "13",
    "R14": "14",
    "R15": "15",
}

symbols = {
    "argument": 'ARG',
    "this": 'THIS',
    "that": 'THAT',
    "local": 'LCL',

}


def get_comment(command, segment, index):
    return f"// {command} {segment} {index}"


def get_push_cmd(index):
    return (
        f"@{index}",
        "D=A",
        "@SP",
        "A=M",
        "M=D",
        "@SP",
        "M=M+1"
    )


def get_pop_cmd(segment: str):
    asm_symbol = symbols.get(segment, "temp")

    return (
        "@SP",
        "AM=M-1",
        "D=M",
        f"@{asm_symbol}",
        "A=M",
        "M=D"
    )


# RAM[SP++] = D -> push the value of D register onto the stack
# @SP, A=M, M=D, @SP, M=M+1->
