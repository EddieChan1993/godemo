/**
	修改map
 */
package main

import "fmt"

func main() {
	m:=make(map[string]int)
	//修改插入元素
	m["Anser"]= 34
	fmt.Println("The value:",m["Anser"])
	//删除元素
	delete(m,"Anser")
	fmt.Println("The value:",m["Anser"])
	//检测某个键存在，键在m中，ok=true
	//不在ok=false，v=0
	v,ok:=m["Anser"]
	fmt.Println("The value:",v,"Present?",ok)
}
