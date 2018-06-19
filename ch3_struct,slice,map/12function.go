package main

import "fmt"

func main() {
	dict := map[string]int{"王五": 60, "张三": 43}
	modify(dict)
	fmt.Println(dict["张三"])

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	modify2(slice)
	fmt.Println(slice)
}

//map传递是改变了的地址
func modify(dict map[string]int) {
	dict["张三"] = 10
}
//切片是3个字段构成的结构类型，
//所以在函数间以值的方式传递的时候，占用的内存非常小，成本很低
//在传递复制切片的时候，其底层数组不会被复制，
//也不会受影响，复制只是复制的切片本身，不涉及底层数组。
func modify2(slice []int) {
	fmt.Printf("%p\n",&slice)
	slice[1]= 10
}