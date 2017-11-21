/**
	在使用:=语法或var=表达式语法
	变量类型由右侧的值推导得出
	常量不能使用:=定义
 */
package main

import "fmt"

func main() {
	i:=42
	f:=float64(i)
	u:=uint(f)

	fmt.Println(f)
	fmt.Println(u)
}
