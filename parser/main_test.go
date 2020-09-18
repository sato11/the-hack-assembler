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
