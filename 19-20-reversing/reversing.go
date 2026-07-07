// 문자열 뒤집기 + 회문검사(회문: 앞으로 읽으나 뒤로 읽으나 같은 문장)
package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reversing(s string) string {
	size := len(s)
	b := make([]byte, size)

	for _, v := range s {
		b[size-1] = byte(v)
		size--
	}
	return string(b)
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
			fmt.Println("exit...")
		}

	}
}
