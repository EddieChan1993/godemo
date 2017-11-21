package main

import (
	"math"
	"fmt"
)

/**
	你可以对包中任意类型定义任意方法，而不仅仅是针对结构体
	但是，不能对来自其他包的类型或基础类型定义方法

	刚刚看到的两个 Abs 方法。一个是在 *Vertex 指针类型上，
	而另一个在 MyFloat 值类型上。 有两个原因需要使用指针接收者。
	首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
	其次，方法可以修改接收者指向的值
 */

type MyFloat float64//定义新的类型MyFloat

//MyFloat实现Abs方法
func(f MyFloat) Abs() float64 {
	if f<0{
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f:=MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

