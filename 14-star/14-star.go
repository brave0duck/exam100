package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func starPrint(n int) {
	s := strings.Repeat("* ", n)

	for i := 0; i < n; i++ {
		fmt.Println(s)
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : star <num>")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Can't convert string")
	}
	starPrint(n)

}
