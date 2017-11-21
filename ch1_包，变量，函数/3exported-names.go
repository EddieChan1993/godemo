/**
	包中，首字母大写的变量或方法，被定义为公共方法，
	导入该包后，可以访问，小写则被定义为私有，只能
	在所在包中访问，其他包无法访问
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
