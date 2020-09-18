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

type compTest struct {
	in  string
	out string
}

func TestComp(t *testing.T) {
	tests := []compTest{
		{"0;JMP", "0101010"},
		{"M=0", "0101010"},
		{"M=1", "0111111"},
		{"M=-1", "0111010"},
		{"A=D", "0001100"},
		{"D;JGT", "0001100"},
		{"D=A", "0110000"},
		{"D=!D", "0001101"},
		{"D=!A", "0110001"},
		{"D=-D", "0001111"},
		{"D=-A", "0110011"},
		{"D=D+1", "0011111"},
		{"D=A+1", "0110111"},
		{"D=D-1", "0001110"},
		{"D=A-1", "0110010"},
		{"D=D+A", "0000010"},
		{"D=D-A", "0010011"},
		{"D=A-D", "0000111"},
		{"D=D&A", "0000000"},
		{"D=D|A", "0010101"},
		{"D=M", "1110000"},
		{"D=!M", "1110001"},
		{"D=-M", "1110011"},
		{"M=M+1", "1110111"},
		{"M=M-1", "1110010"},
		{"D=D+M", "1000010"},
		{"D=D-M", "1010011"},
		{"M=M-D", "1000111"},
		{"D=D&M", "1000000"},
		{"D=D|M", "1010101"},
	}
	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		p.currentCommand = test.in
		comp, err := p.Comp()
		if err != nil {
			t.Errorf("#%d: error returned: %v", i, err.Error())
		}
		if comp != test.out {
			t.Errorf("#%d: got: %v want: %v", i, comp, test.out)
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
