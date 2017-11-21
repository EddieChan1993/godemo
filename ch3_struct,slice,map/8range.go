/**
	range 遍历切片
 */
package main

import "fmt"

func main() {
	var pow=[]int{1,2,54,3,3,12, 23}

	for i,v:=range pow {
		//i下标，v下标所对应元素的拷贝
		fmt.Printf("%d=>%d\n",i,v)
	}

	for i:=range pow {
		fmt.Printf("下标%d:",i)
	}

	for _, v := range pow {
		fmt.Println(v)
	}
}
