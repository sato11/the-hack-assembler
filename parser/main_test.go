package parser

import (
	"bytes"
	"strings"
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

func TestAdvance(t *testing.T) {
	tests := []string{
		"@i",
		"M=1 // i=0",
		"@sum",
		"M=0 // sum=0",
		"(LOOP)",
		"(END)",
	}
	input := strings.Join(tests, "\n")
	b := bytes.NewBufferString(input)
	p := New(b)
	for i, test := range tests {
		err := p.Advance()
		if err != nil {
			t.Errorf("#%d: error returned: %v", i, err.Error())
		}
		if p.currentCommand != test {
			t.Errorf("#%d: got: %v want: %v", i, p.currentCommand, test)
		}
	}
}

type commandTypeTest struct {
	command string
	out     CommandTypes
}

func TestCommandType(t *testing.T) {
	tests := []commandTypeTest{
		{"@i", 0},
		{"@sum", 0},
		{"D=M", 1},
		{"M=M+1", 1},
		{"0;JMP", 1},
		{"(LOOP)", 2},
		{"(END)", 2},
	}
	for i, test := range tests {
		b := bytes.NewBufferString(test.command)
		p := New(b)
		p.currentCommand = test.command
		command, _ := p.CommandType()
		if command != test.out {
			t.Errorf("#%d: got: %v want: %v", i, command, test.out)
		}
	}
}

type symbolTest struct {
	in  string
	out string
}

func TestSymbol(t *testing.T) {
	tests := []symbolTest{
		{"@i", "i"},
		{"@sum", "sum"},
		{"@100", "100"},
		{"(LOOP)", "LOOP"},
		{"(END)", "END"},
	}

	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		p.currentCommand = test.in
		symbol := p.Symbol()
		if symbol != test.out {
			t.Errorf("#%d: got: %v want: %v", i, symbol, test.out)
		}
	}
}

type destTest struct {
	in  string
	out string
}

func TestDest(t *testing.T) {
	tests := []destTest{
		{"0;JMP", "000"},
		{"M=M+1", "001"},
		{"D=M", "010"},
		{"MD=M-1", "011"},
		{"A=A+1", "100"},
		{"AM=A+1", "101"},
		{"AD=A+1", "110"},
		{"AMD=A+1", "111"},
	}
	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		p.currentCommand = test.in
		dest := p.Dest()
		if dest != test.out {
			t.Errorf("#%d: got: %v want: %v", i, dest, test.out)
		}
	}
}

type jumpTest struct {
	in  string
	out string
}

func TestJump(t *testing.T) {
	tests := []jumpTest{
		{"M=1", "000"},
		{"D;JGT", "001"},
		{"D;JEQ", "010"},
		{"D;JGE", "011"},
		{"D;JLT", "100"},
		{"D;JNE", "101"},
		{"D;JLE", "110"},
		{"0;JMP", "111"},
	}
	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		p.currentCommand = test.in
		jump := p.Jump()
		if jump != test.out {
			t.Errorf("#%d: got: %v want: %v", i, jump, test.out)
		}
	}
}
