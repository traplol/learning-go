package main

import (
	"bfi/stack"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var tape [30000]byte
var p = 0
var reader = bufio.NewReader(os.Stdin)
var instructions []*Instruction
var insPtr = 0

const (
	Right   = iota
	Left    = iota
	Add     = iota
	Sub     = iota
	Out     = iota
	In      = iota
	LoopBeg = iota
	LoopEnd = iota
)

type Instruction struct {
	Op     int
	JumpTo int
}

func interp(ins *Instruction) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("p=%d, insPtr=%d\n", p, insPtr)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}()
	switch ins.Op {
	case Right:
		p++
		insPtr++
	case Left:
		p--
		insPtr++
	case Add:
		tape[p]++
		insPtr++
	case Sub:
		tape[p]--
		insPtr++
	case Out:
		fmt.Printf("%c", tape[p])
		insPtr++
	case In:
		tape[p], _ = reader.ReadByte()
		insPtr++
	case LoopBeg:
		if tape[p] != 0 {
			insPtr++
		} else {
			insPtr = ins.JumpTo
		}
	case LoopEnd:
		insPtr = ins.JumpTo
	}
}

func parse(code []byte) {
	for _, c := range code {
		switch c {
		default:
			continue
		case '>':
			instructions = append(instructions, &Instruction{Op: Right, JumpTo: 0})
		case '<':
			instructions = append(instructions, &Instruction{Op: Left, JumpTo: 0})
		case '+':
			instructions = append(instructions, &Instruction{Op: Add, JumpTo: 0})
		case '-':
			instructions = append(instructions, &Instruction{Op: Sub, JumpTo: 0})
		case '.':
			instructions = append(instructions, &Instruction{Op: Out, JumpTo: 0})
		case ',':
			instructions = append(instructions, &Instruction{Op: In, JumpTo: 0})
		case '[':
			instructions = append(instructions, &Instruction{Op: LoopBeg, JumpTo: 0})
		case ']':
			instructions = append(instructions, &Instruction{Op: LoopEnd, JumpTo: 0})
		}
	}
	// Calculate jumps.
	s := stack.NewStack()
	for i := 0; i < len(instructions); i++ {
		op := instructions[i]
		if op.Op != LoopBeg {
			continue
		}
		for j := i; j < len(instructions); j++ {
			op2 := instructions[j]
			if op2.Op == LoopBeg {
				s.Push(1)
			} else if op2.Op == LoopEnd {
				_, empty := s.Pop()
				if empty {
					op.JumpTo = j + 1
					op2.JumpTo = i
					break
				}
			}
		}
	}
}

func run() {
	for insPtr < len(instructions) {
		ins := instructions[insPtr]
		interp(ins)
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s filename\n", args[0])
		return
	}
	code, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "File not found.\n")
		return
	}
	parse(code)
	run()
}
