package parser

import (
	"bufio"
	"io"
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
