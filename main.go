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
	"github.com/sato11/the-hack-assembler/symboltable"
)

// ExitCodeOK and ExitCodeError represent respectively a status code.
const (
	ExitCodeOK int = iota
	ExitCodeError
)

// Client wraps modules and behaves as a uniform interface.
type Client struct {
	parser      *parser.Parser
	code        *code.Code
	symboltable *symboltable.SymbolTable
}

func new(r io.Reader) *Client {
	return &Client{
		parser.New(r),
		code.New(),
		symboltable.New(),
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

func (c *Client) handleFirstPass() error {
	currentAddress := 0
	for c.parser.HasMoreCommands() {
		c.parser.Advance()
		commandType, err := c.parser.CommandType()
		if err != nil {
			return err
		}
		switch commandType {
		case parser.A:
		case parser.C:
			currentAddress++
		case parser.L:
			symbol := c.parser.Symbol()
			c.symboltable.AddEntry(symbol, currentAddress+1)
		}
	}
	return nil
}

func (c *Client) handleSecondPass(buffer *bytes.Buffer) (bytes.Buffer, error) {
	for c.parser.HasMoreCommands() {
		c.parser.Advance()
		commandType, err := c.parser.CommandType()
		if err != nil {
			return *buffer, err
		}
		switch commandType {
		case parser.N:
			// no-op
		case parser.C:
			cInstruction, err := c.handleCInstruction()
			if err != nil {
				return *buffer, err
			}
			buffer.WriteString(fmt.Sprintf("111%s\n", cInstruction))
		case parser.A:
			symbol := c.parser.Symbol()
			address, err := strconv.Atoi(symbol)
			if err != nil {
				address = c.symboltable.GetAddress(symbol)
			}
			aInstruction := fmt.Sprintf("0%015b\n", address)
			buffer.WriteString(aInstruction)
		}
	}
	return *buffer, nil
}

func run(r io.Reader) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	client := new(r)
	err := client.handleFirstPass()
	if err != nil {
		return buffer, err
	}

	client.parser.Reset()

	buffer, err = client.handleSecondPass(&buffer)
	if err != nil {
		return buffer, err
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
	fmt.Printf(output.String())
	os.Exit(ExitCodeOK)
}
