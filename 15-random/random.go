// 15-16 랜덤숫자 생성- 숫자맞추기 게임
// 10번의 찬스. 한번 틀릴때마다 포인트 -10점 감점
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var input int
	var ncount uint8 = 10
	var gamePoint uint8 = 100

	fmt.Println("=== 숫자맞추기 게임 ===")
	fmt.Println("10번의 찬스. 100점 만점. 기회 1번 소모할때마다 -10점")
	fmt.Println("1부터 100사이의 숫자가 랜덤하게 주어집니다.(기회 10번)")

	com := rand.IntN(100) + 1
	
	for ncount > 0 {
		fmt.Print("입력 : ")
		_, err := fmt.Scanf("%d\n", &input)

		if 1 > input || input > 100 {
			fmt.Println("잘못된 입력. 다시 입력하세요")
			continue
		}

		if err != nil {
			fmt.Println("입력한 값은 숫자가 아닙니다.")
			continue
		}

		
		if input == com {
			fmt.Printf("축하합니다. 정답입니다!\n")
			break
		} else {
			ncount -= 1
			gamePoint -= 10
			fmt.Printf("%s. (기회 %d번)\n", check_number(input, com), ncount)
		}
	}
	if ncount > 0 {
		fmt.Printf("%d번만에 정답을 맞추었군요. %d점을 획득했습니다\n", 11-ncount, gamePoint)
	} else {
		fmt.Printf("\n안타깝군요. 정답은 [%d] 이었습니다.\n", com)
	}
	fmt.Println("== 게임을 종료합니다 ==")
}
func check_number(one int, two int) string {
	var message string
	differ := one - two
	if differ < 0 {
		differ = -differ
	}
	if differ > 30 {
		message = "30이상"
	} else if differ > 20 {
		message = "30이하 20초과"
	} else if differ > 10 {
		message = "20이하 10초과"
	} else {
		message = "10이내"
	}
	return message
}
