package main

import (
	"math"
	"fmt"
)
/**
	go没有类，然而，仍然可以在结构体类型上定义方法
	方法接收者，出现在func关键字和方法名之间的参数中
 */

type Vertex struct {
	X,Y float64
}
//v实现Abs方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main() {
	v:=&Vertex{3, 34}
	fmt.Println(v.Abs())
}