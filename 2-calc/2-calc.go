package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Println("Mini Calc ( a + b ) : ")
	fmt.Scanf("%s\n", &input)

	str := strings.Split(input, " ")
	a,err := strconv.Atoi(str[0])
	if err != nil{
		fmt.Println("failed convert string to int")
		return
	}
	b, err := strconv.Atoi(str[2])
	if err != nil{
		fmt.Println("failed convert string to int")
		return
	}
	if result := calc(a, b, str[1]); result != -1 {
		fmt.Println("result : ", result)
	} else {
		fmt.Println("Can't calc")
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
