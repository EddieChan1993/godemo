/**
	指针
 */
package main

import "fmt"

func main() {
	i,j:=53,22323

	p:=&i//生成一个指向i对象地址的指针
	fmt.Println(*p)//通过指针获取底层值

	*p=21//改变指针所指对象的值
	fmt.Println(i)

	p=&j
	*p=*p

	fmt.Println(j)

}
