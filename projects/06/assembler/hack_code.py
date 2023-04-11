from constants import COMP, DEST, JUMP


class HackCode:

    @staticmethod
    def get_binary_dest(dest: str) -> str:
        return DEST[dest]
    
    @staticmethod
    def get_binary_comp(comp: str) -> str:
        return COMP[comp]

    @staticmethod
    def get_binary_jump(jump: str) -> str:
        return JUMP[jump]
