package compiler

type SymbolScope string

const (
	GlobalScope SymbolScope = "GLOBAL"
	LocalScope  SymbolScope = "LOCAL"
)

type SymbolType int

const (
	VarSymbol SymbolType = iota
	FuncSymbol
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
	Type  SymbolType
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
	Outer          *SymbolTable
}

func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func (s *SymbolTable) DefineVar(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Type: VarSymbol}
	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

func (s *SymbolTable) DefineFunc(name string) Symbol {
	symbol := Symbol{Name: name, Index: -1, Type: FuncSymbol, Scope: GlobalScope} // Functions are global
	s.store[name] = symbol
	return symbol
}

func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	if ok {
		return obj, true
	}
	if s.Outer != nil {
		return s.Outer.Resolve(name)
	}
	return obj, false
}
