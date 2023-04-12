from enum import Enum


class InstructionType(str, Enum):
    A = 'A_INSTRUCTION'
    C = 'C_INSTRUCTION'
    L = 'L_INSTRUCTION'


def convert_decimal_to_binary(decimal: int):
    return '{0:016b}'.format(decimal)
