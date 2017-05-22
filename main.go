package main

import (
	"fmt"
)

const (
	iconst = iota
	iadd
	print
	halt
)

type Dyson struct {
	code  []int
	stack []int
	sp    int
	ip    int
}

func (vm *Dyson) alloc(code []int) {
	vm.stack = make([]int, 100)
	vm.sp = -1
	vm.ip = 0
	vm.code = code
}

func (vm *Dyson) exec() {
	for {
		opCode := vm.code[vm.ip]
		vm.ip++
		switch opCode {
		case iconst:
			Iconst(vm)
		case iadd:
			Iadd(vm)
		case print:
			Print(vm)
		case halt:
			return
		}
	}
}

func Iconst(vm *Dyson) {

}

func Iadd(vm *Dyson) {

}

func Print(vm *Dyson) {

}

func main() {
	code := []int{
		iconst, 2,
		iconst, 3,
		iadd,
		halt,
	}
	fmt.Println(code)
}
