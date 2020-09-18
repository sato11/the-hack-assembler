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

// Comp returns the comp mnemonic in the current C-command.
// Should be called only when CommantType() is C.
func (p *Parser) Comp() (string, error) {
	var comp string
	if p.Dest() != "000" {
		comp = strings.Split(p.currentCommand, "=")[1]
	} else {
		comp = strings.Split(p.currentCommand, ";")[0]
	}
	switch comp {
	case "0":
		return "0101010", nil
	case "1":
		return "0111111", nil
	case "-1":
		return "0111010", nil
	case "D":
		return "0001100", nil
	case "A":
		return "0110000", nil
	case "!D":
		return "0001101", nil
	case "!A":
		return "0110001", nil
	case "-D":
		return "0001111", nil
	case "-A":
		return "0110011", nil
	case "D+1":
		return "0011111", nil
	case "A+1":
		return "0110111", nil
	case "D-1":
		return "0001110", nil
	case "A-1":
		return "0110010", nil
	case "D+A":
		return "0000010", nil
	case "D-A":
		return "0010011", nil
	case "A-D":
		return "0000111", nil
	case "D&A":
		return "0000000", nil
	case "D|A":
		return "0010101", nil
	case "M":
		return "1110000", nil
	case "!M":
		return "1110001", nil
	case "-M":
		return "1110011", nil
	case "M+1":
		return "1110111", nil
	case "M-1":
		return "1110010", nil
	case "D+M":
		return "1000010", nil
	case "D-M":
		return "1010011", nil
	case "M-D":
		return "1000111", nil
	case "D&M":
		return "1000000", nil
	case "D|M":
		return "1010101", nil
	default:
		return "", errors.New("invalid comp detected")
	}
}

// Jump returns the jump mnemonic in the current C-command.
// Should be called only when CommandType() is C.
func (p *Parser) Jump() string {
	if strings.Contains(p.currentCommand, ";") {
		jump := strings.Split(p.currentCommand, ";")[1]
		switch jump {
		case "JGT":
			return "001"
		case "JEQ":
			return "010"
		case "JGE":
			return "011"
		case "JLT":
			return "100"
		case "JNE":
			return "101"
		case "JLE":
			return "110"
		case "JMP":
			return "111"
		}
	}
	return "000"
}
