package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string

	fmt.Println("Mini Calc (with spaces. a + b ) : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()

		str := strings.Split(input, " ")
		a, err := strconv.Atoi(str[0])
		if err != nil {
			fmt.Println("failed convert string")
			return
		}
		b, err := strconv.Atoi(str[2])
		if err != nil {
			fmt.Println("failed convert string")
			return
		}
		if result := calc(a, b, str[1]); result != -1 {
			fmt.Println("result : ", result)
		} else {
			fmt.Println("Can't calc")
		}
	}

}

func calc(a int, b int, op string) int {

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return -1
	}
}
