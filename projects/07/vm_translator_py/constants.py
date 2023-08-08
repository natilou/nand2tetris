SEGMENTS = {
    "argument": "ARG",
    "this": "THIS",
    "that": "THAT",
    "local": "LCL",
    "temp": "5",
    "static": "16",
    "pointer": "3",
}


ARITHMETIC_CMDS = {
    "add": "M=D+M",
    "sub": "M=M-D",
    "neg": "M=-M",
    "and": "M=D&M",
    "or": "M=D|M",
    "not": "M=!M",
}

COMPARISON_CMDS = {
    "eq": "D;JEQ",
    "gt": "D;JGT",
    "lt": "D;JLT",
}
