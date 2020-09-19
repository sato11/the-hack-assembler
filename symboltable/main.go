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
