/**
	程序运行的入口时main包
	一个项目只有一个入口
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	aa()
	fmt.Println("My favorite number is",rand.Intn(10))
}