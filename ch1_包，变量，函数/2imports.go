/**
	圆括号组合导入，可同时导入多个包
 */
package main

import (
	"fmt"
	"math"
	_"godemo/test"//导入包，仅仅调用init，通常作为初始化连接之类的
)

func main() {
	fmt.Printf("you have %g problems", math.Sqrt(7))
}
