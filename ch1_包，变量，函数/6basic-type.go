package main

import (
	"math/cmplx"
	"fmt"
)

const name1 string="23"

var (
	name string= "aaa"
	aa byte//uint8的别名
	bb rune//int32的别名
	cc float64
	dd int//首选int

	ToBe bool=false
	MaxInt uint64=1<<64-1
	z complex128=cmplx.Sqrt(-5+12i)
)

func main() {
	const f="%T(%v)\n"
	fmt.Printf(f,ToBe, ToBe)
	fmt.Printf(f,MaxInt, MaxInt)
	fmt.Printf(f,z,z)
	fmt.Printf(f,name,name)
	fmt.Printf(f,aa,aa)
	fmt.Printf(f,bb,bb)
	fmt.Printf(f,cc,cc)
	fmt.Printf(f,dd,dd)
}
