package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/sato11/the-hack-assembler/code"
	"github.com/sato11/the-hack-assembler/parser"
)

// ExitCodeOK and ExitCodeError represent respectively a status code.
const (
	ExitCodeOK int = iota
	ExitCodeError
)

// Client wraps modules and behaves as a uniform interface.
type Client struct {
	parser *parser.Parser
	code   *code.Code
}

func new(r io.Reader) *Client {
	return &Client{
		parser.New(r),
		code.New(),
	}
}

func (c *Client) handleCInstruction() (string, error) {
	dest := c.code.Dest(c.parser.Dest())
	jump := c.code.Jump(c.parser.Jump())
	comp, err := c.code.Comp(c.parser.Comp())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%07b%03b%03b", comp, dest, jump), nil
}

func run(inputReader io.Reader) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	client := new(inputReader)
	for client.parser.HasMoreCommands() {
		client.parser.Advance()
		commandType, err := client.parser.CommandType()
		if err != nil {
			return buffer, err
		}
		switch commandType {
		case parser.N:
			// no-op
		case parser.C:
			cInstruction, err := client.handleCInstruction()
			if err != nil {
				return buffer, err
			}
			buffer.WriteString(fmt.Sprintf("111%s\n", cInstruction))
		case parser.A:
			symbol := client.parser.Symbol()
			address, err := strconv.ParseInt(symbol, 10, 16)
			if err != nil {
				return buffer, err
			}
			aInstruction := fmt.Sprintf("0%015b\n", address)
			buffer.WriteString(aInstruction)
		}
	}
	return buffer, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	output, err := run(reader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(ExitCodeError)
	}
	fmt.Println(output.String())
	os.Exit(ExitCodeOK)
}
