package main

import (
	"strings"
	"fmt"
	"io"
)
/**
	func (T) Read(b []byte) (n int, err error)
	Read 用数据填充指定的字节 slice，
	并且返回填充的字节数和错误信息。 在遇到数据流结尾时，
	返回 io.EOF 错误
 */
func main() {
	//创建一个strings.Reader
	r:=strings.NewReader("Hello,Reader!")

	//每次8字节的速度读取它的输出
	b:=make([]byte, 8)

	for  {
		n,err:=r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n",n,err,b)
		fmt.Printf("b[:n]=%q\n",b[:n])
		if err ==io.EOF{
			fmt.Println("没有内容")
			break
		}
	}
}