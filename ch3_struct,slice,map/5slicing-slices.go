/**
	对slice切片
 */
package main

import "fmt"

func main() {
	//0~len(s)
	s:=[]int{2,3,6,8,54, 2}

	fmt.Println("s==", s)
	fmt.Println("s[1:4]==",s[1:4])//含前端，不包含后端

	fmt.Println("s[:3]==",s[:3])
	fmt.Println("s[4:]==",s[4:])
}
