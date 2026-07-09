// 단어갯수 세기
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("문장에서 단어갯수를 리턴합니다. 문장을 입력해보세요")
	if scanner.Scan() {
		s := scanner.Text()

		if n := wordCount(s); n > 0 {
			fmt.Printf("%d개의 단어로 이루어져있습니다\n",n)
		} else {
			fmt.Println("단어가 없습니다")
		}
	}
}

// 문장에서 단어 갯수 리턴 
//  []string = strings.Fields()함수는 공백문자로 구분, 이스케이프시퀀스포함
func wordCount(s string) int {
	wl := strings.Fields(s)
	return len(wl)
}
