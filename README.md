# The Hack Assembler
A golang implementation of `The Hack Assembler` based on _The Elements of Computing Systems: Building a Modern Computer from First Principles_ aka [Nand2Tetris](https://www.nand2tetris.org/).

## Modules
The package is made up of three modules, namely:
- parser, which encapsulates access to the input code and provides convenient access to it.
- code, which translates hack assembly language mnemonics into binary code.
- symboltable, which creates and maintains the correspondence between symbols and their RAM/ROM addresses.

However, the `main.go` has already made use of these modules so you do not have to mind the implementation details.

## Installation
Simply clone the repository and run `go install`.

## How to use
Source assembly files are provided under `./testdata` directory. Use them to see `the-hack-assembler` translates assembly into binary code.

The executable expects STDIN to provide source lines. Pass the file contents via command line like below.
```
the-hack-assembler < testdata/Add.asm
```

Likewise when you want to save the output to as a file, redirect it manually. The result files are conventionally given extension `.hack`.
```
the-hack-assembler < testdata/Max.asm > Max.hack
```
