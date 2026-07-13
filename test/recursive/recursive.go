// 1부터 10까지의 숫자를 재귀로 출력하세요
package main

import (
	"fmt"
)

func main() {
	printNum2(10)
}
// 10 9 8... 1
func printNum(n int8) int8 {
	if n <= 0 {
		return 0
	}
	fmt.Println(n)
	return printNum(n - 1)

}
// 1 2 3 ... 10
func printNum2(n int8) int8{
	if n > 0 && n < 20{
		fmt.Println(n-9)
		return printNum2(n+1)
	}
	return 0
}