package main

import (
	"fmt"
)

func fac_v1(n int) int {
	if n <= 1 {
		return n
	}
	fmt.Println(n)
	return n * fac_v1(n-1)
}
func fac_v2(n int) int {
	var result int = 1
	for i := n; i > 0; i-- {
		result *= i
		fmt.Println(result)
	}
	return result
}
func main() {
	var input int
	fmt.Println("팩토리얼 계산기")
	fmt.Print("입력 : ")
	fmt.Scanf("%d\n", &input)

	fmt.Println("재귀 방식 : ", fac_v1(input))
	fmt.Println("for문 방식 : ", fac_v2(input))
}
