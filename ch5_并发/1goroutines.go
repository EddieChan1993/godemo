package main

import (
	"time"
	"fmt"
)
/**
	gorountine在相同的地址空间中运行，因此访问共享内存
	必须进行同步
 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000*time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")

	go func(aa string) {
		//内置函数
	}("a")
}
