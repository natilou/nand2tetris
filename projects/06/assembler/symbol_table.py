from constants import PREDEFINED_R, PREDEFINED_SYMBOLS


class SymbolTable:

    def __init__(self) -> None:
        self.table = {
            **PREDEFINED_R,
            **PREDEFINED_SYMBOLS,
        }

    def add_entry(self, symbol: str, address: str) -> None:
        self.table[symbol] = address
    
    def contains(self, symbol: str) -> bool:
        return symbol in self.table
    
    def get_address(self, symbol: str) -> str:
        return self.table[symbol] if symbol in self.table else 0
