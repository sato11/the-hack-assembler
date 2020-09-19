package symboltable

// SymbolTable wraps resolution of symbols into address.
type SymbolTable struct {
	table map[string]int
}

// New initializes an empty hash table.
func New() *SymbolTable {
	return &SymbolTable{
		make(map[string]int),
	}
}

// AddEntry adds the pair to the table.
func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.table[symbol] = address
}

// Contains returns true if the symbol table contain the given symbol.
func (s *SymbolTable) Contains(symbol string) bool {
	return s.table[symbol] > 0
}

// GetAddress returns the address associated with the symbol.
func (s *SymbolTable) GetAddress(symbol string) int {
	return s.table[symbol]
}
