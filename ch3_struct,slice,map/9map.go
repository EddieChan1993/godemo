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
//声明但是没有初始化分配内存空间，无法使用
var m map[string]Vertex

//声明并初始化
var m1=map[string]Vertex{
	"Tom":Vertex{23,34},
}

func main()  {
	m=make(map[string]Vertex)//初始化，分配了空间，可以使用
	m["Bell Labs"]=Vertex{
		40.34234,-123.2312,
	}
	fmt.Print(m["Bell Labs"])

	m1["Jack"]=Vertex{
		40.34234,-123.2312,
	}
	fmt.Print(m1["Tom"])
	fmt.Print(m1["Jack"])

	m2:=make(map[string]Vertex)//声明并初始化一个空的map
	fmt.Println(m2)
	m2["Jack"]=Vertex{
		40.34234,-123.2312,
	}
	fmt.Print(m2["Jack"])
}


