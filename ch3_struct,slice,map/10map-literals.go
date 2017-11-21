/**
	map文法
 */
package main

import "fmt"

type Vertex struct {
	Lat,Long float64
}


var m=map[string]Vertex{
	"Bell":Vertex{40.234,-324},
	"Google":Vertex{32.23423,-123.123},
}

var n=map[string]Vertex{
	"aa":{123,2322},
	"bb":{3434,333},
}

func main() {
	fmt.Println(m)
	fmt.Println(n)
}
