package core

import (
	"fmt"
)

type Instruction byte

const (
	InstrPush Instruction = 0x0a // 10
	InstrAdd  Instruction = 0x0b // 1
	InstrByte Instruction = 0x0c
	InstrPack Instruction = 0x0d
	InstrSub  Instruction = 0x0e
)

type Stack struct {
	data []any
	sp   int
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]any, size),
		sp:   0,
	}
}

func (s *Stack) Push(v any) {
	s.data[s.sp] = v
	s.sp++
}

func (s *Stack) Pop() any {
	value := s.data[0]
	s.data = append(s.data[:0], s.data[1:]...)
	s.sp--

	return value
}

type VM struct {
	data  []byte
	ip    int
	stack Stack
}

func NewVM(data []byte) *VM {
	return &VM{
		data:  data,
		ip:    0,
		stack: *NewStack(128),
	}
}

func (vm *VM) Run() error {
	for {
		instr := Instruction(vm.data[vm.ip])

		fmt.Println(instr)

		if err := vm.Exec(instr); err != nil {
			return err
		}

		vm.ip++

		if vm.ip > len(vm.data)-1 {
			break
		}
	}

	return nil
}
func (vm *VM) Exec(instr Instruction) error {
	switch instr {
	case InstrPush:
		vm.stack.Push(int(vm.data[vm.ip-1]))
	case InstrByte:
		vm.stack.Push(byte(vm.data[vm.ip-1]))
	case InstrPack:
		n := vm.stack.Pop().(int)
		b := make([]byte, n)

		for i := 0; i < n; i++ {
			b[i] = vm.stack.Pop().(byte)
		}

		vm.stack.Push(b)

	case InstrAdd:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a + b
		vm.stack.Push(c)

	case InstrSub:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a - b
		vm.stack.Push(c)
	}

	return nil
}
