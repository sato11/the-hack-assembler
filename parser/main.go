package parser

import (
	"bufio"
	"io"
)

// Parser wraps the input stream to which operations are intended
type Parser struct {
	reader *bufio.Reader
}

// New converts input stream to Parser
func New(input io.Reader) *Parser {
	return &Parser{
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
