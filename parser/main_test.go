package parser

import (
	"bytes"
	"testing"
)

type hasMoreCommandsTest struct {
	in  string
	out bool
}

func TestHasMoreCommands(t *testing.T) {
	tests := []hasMoreCommandsTest{
		{"", false},
		{"\n", true},
		{"@value", true},
		{"M=1", true},
		{"(LOOP)", true},
		{"D=D-A", true},
		{"0;JMP // infinite loop", true},
	}
	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		out := p.HasMoreCommands()
		if out != test.out {
			t.Errorf("#%d: input: %v, got: %v want: %v", i, test.in, out, test.out)
		}
	}
}
