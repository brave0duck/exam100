// 아나그램 검사?
// 두 문장이 알파벳 순서만 다를 뿐,
// 1. 구성하는 문자의 종류와 2. 개수가 정확히 일치하는지 확인하는 검사입니다.
// listen -> eilnst, silent -> eilnst
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("checker anagram")
	fmt.Println("input two string and check")

	var s1 string
	var s2 string
	fmt.Print("input : ")
	if scanner.Scan() {
		s1 = scanner.Text()
	}
	fmt.Print("input second : ")
	if scanner.Scan() {
		s2 = scanner.Text()
	}

	s1 = strings.Join(strings.Fields(s1), "")
	s2 = strings.Join(strings.Fields(s2), "")
	if checkAnagram(s1, s2) {
		fmt.Println("anagram")
	} else {
		fmt.Println("not anagram")
	}

}
func checkAnagram(left string, right string) bool {
	var isAna bool
	s1 := strings.ToLower(left)
	s2 := strings.ToLower(right)

	for _, v1 := range s1 {
		isAna = false
		for _, v2 := range s2 {
			if v1 == v2 {
				isAna = true
				break
			}
		}
		if isAna == false {
			return false
		}
	}
	return isAna
}
