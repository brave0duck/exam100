package main

import(
	"fmt"
	"math/rand/v2"
)
func dicing(n int) []int{
	dice := map[int]int{}
		
	for i:=0 ; i < n ; i++ {
		dice[rand.IntN(6)+1]++
	}
}
func main(){
	var usr int
	fmt.Println("주사위 시뮬레이터")
	fmt.Println("몇번을 주사위 돌려볼까요? ")

	fmt.Scanf("%d\n", &usr)

	a := dicing(usr)
	fmt.Println("결과 : ")
	for k,v:=range a{
		fmt.Printf("[%d] : %d", k,v)	
	}
}
