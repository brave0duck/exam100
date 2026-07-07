// 주사위 시뮬레이터
package main

import (
	"fmt"
	"math/rand/v2"
)

func dicing(n int) []int {
	dice := make([]int, 6)

	for i := 0; i < n; i++ {
		dice[rand.IntN(6)]++
	}
	return dice
}
func main() {
	var usr int
	fmt.Println("주사위 시뮬레이터")
	fmt.Println("횟수 : ")

	fmt.Scanf("%d\n", &usr)

	a := dicing(usr)
	fmt.Println("--> 결과 (기대값 : 0.1666)")
	for k, v := range a {
		fmt.Printf("[%d] : %d/%d (%.2f)\n", k+1, v, usr, float64(v)/float64(usr))
	}
}
