package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Println("윤년계산 프로그램")
	fmt.Print("연도 : ")
	fmt.Scanf("%d\n", &year)

	if isYunNeon(year) {
		fmt.Printf("%d년 --> [윤년]입니다",year)
	} else {
		fmt.Printf("%d년 --> [평년]입니다.",year)
	}
}
// 윤년 공식
// 1. 4로 나누어 떨어지면 윤년
// 2.100으로 나누어 떨어지면 평년
// 3. 100으로 나누어 떨어지나 400으로 나누어 떨어지면 윤년

func isYunNeon(year int) bool {
	var result bool = false
	if year%4 == 0 {
		result = true
	}
	if year%100 == 0 {
		if year%400 == 0 {
			result = true
		} else {
			result = false
		}
	}
	return result
}
