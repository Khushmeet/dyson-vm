package main

import (
	"fmt"
)

const (
	iconst = iota
	iadd
	jmplt
	print
	halt
)

type Opcode struct {
	name     string
	num_args int
}

var opMetadata = map[int]Opcode{
	iconst: Opcode{"ICONST", 1},
	iadd:   Opcode{"IADD", 0},
	jmplt:  Opcode{"JMPLT", 2},
	print:  Opcode{"PRINT", 0},
	halt:   Opcode{"HALT", 0},
}

type Dyson struct {
	code  []int
	stack []int
	sp    int
	ip    int
}

func (vm *Dyson) trace() {
	address := vm.ip
	op := opMetadata[vm.code[vm.ip]]
	stack := vm.stack[0:vm.sp+1]
	fmt.Printf("%04d: %s \t%v\n", address, op.name, stack)
}

func (vm *Dyson) alloc(code []int) {
	vm.stack = make([]int, 100)
	vm.sp = -1
	vm.ip = 0
	vm.code = code
}

func (vm *Dyson) exec() {
	for {
		vm.trace()
		opCode := vm.code[vm.ip]
		vm.ip++
		switch opCode {
		case iconst:
			Iconst(vm)
		case iadd:
			Iadd(vm)
		case jmplt:
			Jmplt(vm)
		case print:
			Print(vm)
		case halt:
			return
		}
	}
}

func Iconst(vm *Dyson) {
	value := vm.code[vm.ip]
	vm.ip++
	vm.sp++
	vm.stack[vm.sp] = value
}

func Iadd(vm *Dyson) {
	top := vm.stack[vm.sp]
	vm.sp--
	second := vm.stack[vm.sp]
	vm.stack[vm.sp] = top + second
}

func Jmplt(vm *Dyson) {
	value := vm.code[vm.ip]
	vm.ip++
	address := vm.code[vm.ip]

	if value > vm.stack[vm.sp] {
		vm.ip = address
	}
}

func Print(vm *Dyson) {
	top := vm.stack[vm.sp]
	vm.sp--
	fmt.Println(top)
}

func main() {
	code := []int{
		iconst, 2,
		iconst, 3,
		iadd,
		jmplt, 10, 2,
		print,
		halt,
	}

	dysonVM := &Dyson{}
	dysonVM.alloc(code)
	dysonVM.exec()
}
