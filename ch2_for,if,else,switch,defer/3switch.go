package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")

	switch os:=runtime.GOOS;os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.",os)

	}


	fmt.Println("\nWhen's Saturday?")
	today:=time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today+0:
		fmt.Println("Today.")
	case today+1:
		fmt.Println("Tomorrow")
	case today+2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	t:=time.Now()
	fmt.Println(t.Hour())
	switch  {
	case t.Hour()< 12:
		fmt.Println("Good morning")
	case t.Hour()<17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")

	}
}
