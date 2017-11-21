/**
	append插入切片元素
 */
package main

import "fmt"

func main() {
	var a []int
	printSlice("a",a)

	a=append(a, 0)
	printSlice("a", a)

	a=append(a, 12)
	printSlice("a", a)

	a=append(a,23, 54)
	printSlice("a",a)
}

func printSlice(s string,x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}