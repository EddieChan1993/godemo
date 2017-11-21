/**
	map映射键到值
	map使用前必须用make创建
	值为nil的map是空的，并且不能对其赋值
 */
package main

import "fmt"

type Vertex struct {
	Lat,Long float64
}

var m map[string]Vertex//定义,如果定义，没有赋值则，必须要make，才可以使用map操作

func main()  {
	m=make(map[string]Vertex)//创建，可以不需要定义，直接创建map
	m["Bell Labs"]=Vertex{
		40.34234,-123.2312,
	}
	fmt.Print(m["Bell Labs"])

}


