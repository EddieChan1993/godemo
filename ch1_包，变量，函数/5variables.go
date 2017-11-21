package main

import "fmt"

var c,python,java bool//变量声明

var i,j int=1,2//初始化变量

func main() {
	var i int
	var c1,python2,java2=true,false,"no!"
	k:=3//短声明变量，只能用在函数内部

	fmt.Println(i,c,python,java)
	fmt.Println(i,j,c1,python2,java2)
	fmt.Println(k)
}
