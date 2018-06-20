package main

import (
	"sync"
	"godemo/ch7_高阶/pool/com"
	"io"
	"sync/atomic"
)

const (
	maxGoroutine=5
	//资源池的大小
	poolRes=2
)

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)

	p,err:=com.New(createConnection,poolRes)
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {

		}(query)
	}
}

//生成数据库连接的方法，以供资源池使用
func createConnection() (io.Closer, error) {
	//并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}