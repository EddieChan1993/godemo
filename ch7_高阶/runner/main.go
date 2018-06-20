package main

import (
	"time"
	"log"
	"godemo/ch7_高阶/runner/com"
	"os"
)

func main() {
	timeout:=12* time.Second
	r:=com.New(timeout)

	for i := 0; i < 5; i++ {
		r.Add(createTask(),createTask())
	}
	if err := r.Start(); err != nil {
		switch err {
		case com.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case com.ErrInterrrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("doing%d", id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}