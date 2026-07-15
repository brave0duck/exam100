package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)
// 재귀방식
func fac_v1(n int) int {
	if n <= 1 {
		return n
	}
	return n * fac_v1(n-1)
}
// int64이상의 숫자도 가능한 팩토리얼 함수. math.big사용
func fac_v2(n int) *big.Int {
	f := big.NewInt(1)

	for i := n; i > 0; i-- {
		f.Mul(f, big.NewInt(int64(i)))
	}
	return f
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("사용법: factorial <숫자>")
		return
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("사용법: factorial <숫자>")
		return
	}

	fmt.Println("재귀 방식 : ", fac_v1(input))
	fmt.Println("for문 방식 : ", fac_v2(input))

}
