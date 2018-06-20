package main

import (
	"sync"
	"fmt"
	"sync/atomic"
)

/**
	go build -race
	数据竞争查看
 */
var(
	count int32
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

//原子操作函数
func incCount() {
	defer wg.Done()
	atomic.StoreInt32(&count,atomic.LoadInt32(&count))
}

//互斥锁
func incCount2() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value:= count
		value++
		count= value
		mutex.Unlock()
	}
}

//通道