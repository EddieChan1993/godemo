/**
	定义切片
 */
package main
import (
	"fmt"
	"strings"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("s==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]==%d\n", i, s[i])
	}

	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	game[0][2]="X"
	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n",strings.Join(s[i]," "))
	}
}
