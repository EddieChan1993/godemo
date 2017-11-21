/**
	变量在定义时没有明确的初始化时会赋值为零值
	数值类型:0
	布尔类型:false
	字符串:""
 */
package main

import "fmt"

func main() {
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("%v %v %v %q\n",i,f,b,s)
}
