/**
	构造slice
 */
package main

import "fmt"

func main() {
	var z []int//[],0,0，z为nil,声明和初始化了一个空的切片，长度和容量都为0
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	fmt.Println(z,len(z),cap(z))//
	if z ==nil {
		fmt.Println("nil")
	}

	a:=make([]int,5)//创建值为[0 0 0 0 0]
	printSlice("a",a)

	b:=make([]int,0, 5)//创建值为[],0,5,不为nil

	if b==nil {
		fmt.Println("nil")
	}
	printSlice("b",b)

	c:=b[:2]//长度改变，容量不变
	printSlice("c", c)

	d:=c[2:5]//长度和容量都改变
	printSlice("d",d)

}

func printSlice(s string ,x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}