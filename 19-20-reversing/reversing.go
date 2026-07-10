// 문자열 뒤집기 + 회문검사(회문: 앞으로 읽으나 뒤로 읽으나 같은 문장)
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func Reversing(s string) string {
	size := utf8.RuneCountInString(s) // 유니코드의 글자길이를 구하는 함수
	r := make([]rune, size)

	for _, v := range s {
		r[size-1] = rune(v)
		size--
	}
	return string(r)
}
func isPalindrome(s string) bool {

	r := Reversing(s)
	if r == s && len(s) > 1 {
		return true
	}
	return false
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("reversing string and check Palindrome")
	fmt.Print("input string : ")
	if scanner.Scan() {
		input := scanner.Text()
		if len(input) > 0 {
			re := Reversing(input)
			fmt.Println("Input : ", input)
			fmt.Println("Reverse : ", re)

			if isPalindrome(input) {
				fmt.Println("---> This is Palindrome !")
			} else {
				fmt.Println("---> Not Palindrome :)")
			}
		}else{
			fmt.Println("word too short")
		}

	}
}
