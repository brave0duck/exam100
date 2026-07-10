// 리스트 중복제거. 리스트를 맵셋으로 변환. 그후 다시 리스트로 변환
package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func convertSet(lst []int) []int {
	var m  = map[int]bool{} // {}이 없으면 nil이 되어 오류발생
	var result = []int{}

	for _, v := range lst {
		if !m[v] {
			m[v] = true
		}
	}
	for key, val := range m {
		if val == true {
			result = append(result, key)
		}
	}
	return result
}
func main() {
	var input string
	var sample = make([]int, 20)

	fmt.Println("리스트 중복제거. 리스트중 중복되는 멤버를 모두 제거함")
LOOP:
	for {
		fmt.Println("계속 하시겠습니까? (y/n) : ")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "y", "Y":
			fmt.Println("크기20인 랜덤리스트를 만듭니다")

			for i := 0; i < 20; i++ {
				sample[i] = rand.IntN(20) + 1
			}
			slices.Sort(sample)
			fmt.Println(sample)
			fmt.Println("리스트 중복제거를 합니다")
			conv := convertSet(sample)
			slices.Sort(conv)
			fmt.Println(conv)

		case "n", "N":
			break LOOP
		default:
			fmt.Println("잘못된 입력입니다")

		}
	}

}
