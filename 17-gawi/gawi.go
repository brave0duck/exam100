// 가위-바위-보 게임 구현
package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func main() {
	game := []string{"가위", "바위", "보"}

	fmt.Println("===== 가위-바위-보 게임 ====")
	fmt.Println("Coumputer가 가위.바위.보 중 하나를 선택했습니다")

	com_i := rand.IntN(3)
	com := game[com_i]

	// usr input
	var usr string
	var usr_i int
	fmt.Print("당신의 선택은 (가위,바위,보)? ")
	fmt.Scanf("%s\n", &usr)

	if slices.Contains(game, usr) {
		for k, v := range game {
			if v == usr {
				usr_i = k
			}
		}
	} else {
		fmt.Println("잘못된 입력.")
		return
	}

	if usr_i > com_i {
		if usr_i == 2 && com_i == 0 {
			fmt.Printf("==> 졌습니다. You:%s, Com:%s\n", usr, com)
		} else {
			fmt.Printf("==> 이겼습니다! You : %s, Com : %s\n", usr, com)
		}

	} else if usr_i < com_i {
		if usr_i == 0 && com_i == 2 {
			fmt.Printf("==> 이겼습니다! You:%s, Com:%s\n", usr, com)
		} else {
			fmt.Printf("==> 졌습니다.  You: %s, Com:%s\n", usr, com)
		}
	} else {
		fmt.Printf("==> 비겼습니다. You:%s, Com:%s\n", usr, com)
	}

}
