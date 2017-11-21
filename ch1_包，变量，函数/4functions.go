/**
	函数
 */
package main

import (
	"fmt"
)

func add(x int,y int) int{
	return x+y
}

func add2(x, y int) (int,int) {
	return x+y,x-y
}

func split(sum int) (x, y int) {
	x=sum*4/ 9
	y=sum- x
	return
}

func show() float64{
	return 12.23
}

func main() {
	fmt.Println(add(42,13))
	fmt.Println(add2(23,234))
	fmt.Println(split(200))
	fmt.Println(show())
}
