package main

import (
	"fmt"
	"os"
	"strconv"
)

func gugudan(n int) {
	if n >= 2 && n <= 9 {
		for i := 1; i <= 9; i++ {
			fmt.Printf("%2d x %2d = %3d\n", n, i, n*i)
		}
	}else{
		fmt.Println("not implement. 2-9 only")
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : gugudan <number>")
		return
	}
	p, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Can't convert", p)
		return
	}
	gugudan(p)

}
