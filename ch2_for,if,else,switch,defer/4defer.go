/**
	defer延迟调用的参数会立刻生成，
	语句会延迟函数的执行到上层函数返回

	延迟函数调用被压入一个栈中。
	当函数返回的时候，会按照后进先出的顺序
	调用被延迟的函数
 */
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
