// 7. 홀수/짝수 판별
// 8. 최대값/최소값 찾기
package main

import (
	"fmt"
	"math/rand"
)

func isOdd(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}
func minMax(n []int) (int,int) {
	if n == nil {
		return -1, -1
	}
	min := n[0]
	max := n[0]

	for _, v := range n {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}
func main() {
	nlist := []int{}
	fmt.Println("홀수/짝수 and 최대값/최소값 찾기")
	fmt.Println("리스트를 랜덤생성중 ...")
	fmt.Println("완료!")

	for i := 0; i < 20; i++ {
		nlist = append(nlist, rand.Intn(100)+1)
	}

	fmt.Println("생성 결과 : ", nlist)
	fmt.Println("------------------------------------")
	for i := 0; i < 20; i++ {
		if isOdd(nlist[i]) {
			fmt.Printf("%d는 홀수\n", nlist[i])
		} else {
			fmt.Printf("%d는 짝수\n", nlist[i])
		}
	}
	max, min := minMax(nlist)
	fmt.Printf("리스트의 최대값/ 최소값은 : %d, %d 입니다\n", max, min)
}
