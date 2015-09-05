package main

import (
	"bfi/stack"
	"bufio"
	"fmt"
	"os"
)

var tape [30000]byte
var p = 0
var reader = bufio.NewReader(os.Stdin)
var jmpStack = stack.NewStack()
var instructions []int
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

type Instruction int

func interp(ins Instruction) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("p=%d, insPtr=%d\n", p, insPtr)
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}()
	switch ins {
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
		fmt.Print(tape[p])
		insPtr++
	case In:
		tape[p], _ = reader.ReadByte()
		insPtr++
	case LoopBeg:
		insPtr++
		fmt.Println("Pushing IP:", insPtr)
		jmpStack.Push(insPtr)
	case LoopEnd:
		fmt.Printf("Stack: %s\n", jmpStack.String())
		insPtr, _ = jmpStack.Pop()
		fmt.Println("Popping IP:", insPtr)
	}
}

func parse(code string) {
	for _, c := range code {
		switch c {
		default:
			continue
		case '>':
			instructions = append(instructions, Right)
		case '<':
			instructions = append(instructions, Left)
		case '+':
			instructions = append(instructions, Add)
		case '-':
			instructions = append(instructions, Sub)
		case '.':
			instructions = append(instructions, Out)
		case ',':
			instructions = append(instructions, In)
		case '[':
			instructions = append(instructions, LoopBeg)
		case ']':
			instructions = append(instructions, LoopEnd)
		}
	}
}

func run() {
	for insPtr < len(instructions) {
		ins := Instruction(instructions[insPtr])
		interp(ins)
	}
}

func main() {
	code := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	parse(code)
	run()
}
