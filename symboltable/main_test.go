package symboltable

import "testing"

type addEntryTest struct {
	symbol  string
	address int
	total   int
}

func TestAddEntry(t *testing.T) {
	tests := []addEntryTest{
		{"i", 1, 1},
		{"sum", 2, 2},
		{"i", 1, 2},
		{"LOOP", 3, 3},
		{"sum", 2, 3},
		{"END", 4, 4},
	}
	s := New()
	for i, test := range tests {
		s.AddEntry(test.symbol, test.address)
		if len(s.table) != test.total {
			t.Errorf("#%d: got: %v wanted: %v", i, len(s.table), test.total)
		}
	}
}

type containsTest struct {
	in  string
	out bool
}

func TestContains(t *testing.T) {
	tests := []containsTest{
		{"contained", true},
		{"notcontained", false},
		{"included", true},
		{"notincluded", false},
		{"", false},
	}
	s := New()
	s.table["contained"] = 1
	s.table["included"] = 3
	for i, test := range tests {
		if s.Contains(test.in) != test.out {
			t.Errorf("#%d: got: %v wanted: %v", i, test.in, test.out)
		}
	}
}
