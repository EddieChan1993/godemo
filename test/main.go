package main

import (
	"github.com/panjf2000/ants"
	"fmt"
)

func main() {
	ants.NewPool(1)
}

func demoFunc()  {
	fmt.Println("Hi")
}