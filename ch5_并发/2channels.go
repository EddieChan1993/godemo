package main

import "fmt"
/**
默认情况下，在另一端准备好之前，发送和接收都会阻塞。直到一方完成
可用于抢票，高并发下
这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步
 */

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum//将和送入c
}

func main() {
	a := []int{7, 2, 8, -8, 3, 2}

	c := make(chan int)//和map与slice一样，使用前必须make

	go sum(a[:len(a) / 2], c)
	go sum(a[len(a) / 2:], c)

	x,y:=<-c,<-c//从c中获取

	fmt.Println(x,y,x+y)
}


