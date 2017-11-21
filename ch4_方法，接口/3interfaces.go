package main

import (
	"math"
	"fmt"
)
/**
	接口类型是由一组方法定义的集合
	接口类型的值可以存放实现这些方法的任何值
 */
//定义Abser接口
//该接口是Abs方法的集合
type Abser interface {
	Abs() float64
}
type MyFloat float64
//MyFloat是Abs方法
func (f MyFloat) Abs()float64  {
	if f<0{
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X,Y float64
}
//&Vertex实现Abs方法
func (v *Vertex)Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main()  {
	var a Abser
	f:=MyFloat(-math.Sqrt2)
	v:=Vertex{3,5}

	a=f
	a=&v
	//a=v,Verte没有实现Abs方法
	fmt.Println(a.Abs())
}
