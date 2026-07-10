// 중복문자 제거 
package main

import (
	"fmt"
)

func delChar(s string) string {
	var check rune
	var result []rune

	for _, v := range s {
		if check != rune(v) {
			check = rune(v)
			result = append(result, check)
		}
	}
	return string(result)
}
func main() {
	var sample1 string = "aabcddeff" // ->abcdef
	
	s := delChar(sample1)
	fmt.Println(sample1, "-->", s)
}
