package parser

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Parser wraps the input stream to which operations are intended
type Parser struct {
	currentCommand string
	reader         *bufio.Reader
}

// New converts input stream to Parser
func New(input io.Reader) *Parser {
	return &Parser{
		"",
		bufio.NewReader(input),
	}
}

// HasMoreCommands returns true when there are more commands in the input
func (p *Parser) HasMoreCommands() bool {
	_, err := p.reader.Peek(1) // returns err on EOF
	if err != nil {
		return false
	}
	return true
}

// Advance reads the next line from the input and makes it the current command.
// Should be called only if `HasMoreCommands()` is true.
func (p *Parser) Advance() error {
	b, _, err := p.reader.ReadLine()
	if err != nil {
		return err
	}
	p.currentCommand = string(b)
	return nil
}

// CommandTypes represent the return value for func CommandType()
type CommandTypes int

const (
	// A represents A-instruction
	A CommandTypes = iota
	// C represents C-instruction
	C
	// L represents pseudo-command in the form of (Xxx)
	L
	// E represents exception: never returned without error
	E
)

// CommandType returns the type of the current command
func (p *Parser) CommandType() (CommandTypes, error) {
	c := p.currentCommand
	if strings.HasPrefix(c, "@") {
		return A, nil
	}
	if strings.Contains(c, "=") {
		return C, nil
	}
	if strings.Contains(c, ";") {
		return C, nil
	}
	if strings.HasPrefix(c, "(") && strings.HasSuffix(c, ")") {
		return L, nil
	}
	return E, errors.New("invalid command detected")
}

// Symbol returns the symbol or decimal Xxx of the current command @Xxx or (Xxx).
// Should be called only when CommandType() is A or L.
func (p *Parser) Symbol() string {
	return strings.Trim(p.currentCommand, "@()")
}

// Dest returns the dest mnemonic in the current C-command.
// Should be called only when CommandType() is C.
func (p *Parser) Dest() string {
	dest := strings.Split(p.currentCommand, "=")[0]
	switch dest {
	case "M":
		return "001"
	case "D":
		return "010"
	case "MD":
		return "011"
	case "A":
		return "100"
	case "AM":
		return "101"
	case "AD":
		return "110"
	case "AMD":
		return "111"
	default:
		return "000"
	}
}
