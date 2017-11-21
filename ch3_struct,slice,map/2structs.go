/**
	结构体
 */
package main

import "fmt"
//结构体
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1,2})

	//结构体字段
	v:=Vertex{1, 2}
	v.X=23
	fmt.Println(v)

	//结构体指针
	p:=&v
	p.X=1e9
	fmt.Println(v)

	//结构体文法
	v1:=Vertex{1, 2}//类型为Vertex
	v2:=Vertex{X:23}//Y:0被忽略
	v3:=Vertex{}//X:0和Y:0
	p4:=&Vertex{1, 3}//类型为*Vertex

	fmt.Println(v1,v2,v3,p4)
}
