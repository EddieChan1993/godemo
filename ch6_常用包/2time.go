package main

import (
	"fmt"
	"time"
)

func main() {
	//时间戳
	timeStampInt64:=time.Now().Unix()
	fmt.Println(timeStampInt64)

	//时间戳->标准格式
	tm :=time.Unix(timeStampInt64, 0)
	fmt.Println(tm.Format("2006-01-02 03:04:05"))
	fmt.Println(tm.Format("02/01/2006 15:04:05"))

	//标准格式->时间戳
	tm2, _:=time.Parse("2006-01-02 03:04:05",tm.Format("2006-01-02 03:04:05"))
	fmt.Println(tm2.Unix())
}