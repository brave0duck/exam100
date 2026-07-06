// 피보나치 수열
package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibonacci(n int) []int {
	f := []int{1, 1}
	for i := 2; i < n; i++ {
		f = append(f, f[i-1]+f[i-2])
	}
	return f
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage : fibonacci <number>\n")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		fmt.Println("Can't execute. parameter problome")
		return
	}

	switch num {
	case 1:
		fmt.Println("[1]")
	case 2:
		fmt.Println("[1,1]")
	default:
		f := fibonacci(num)
		fmt.Println(f)
		fmt.Printf("%dth's fibonacci : [%d]\n", num, f[len(f)-1])
	}
}
