// 가장 많이 등장한 단어들 찾기
// 조바이든 2022년 대통령취임 연설문으로 예제작성
// 1. those - must - you - america - more - story 
package main

import (
	"fmt"
	"os"
	"strings"
)

func file_read(file_name string) string {
	f, err := os.OpenFile(file_name, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer f.Close()

	f_stat, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	data := make([]byte, f_stat.Size())
	n, err := f.Read(data)
	if err != nil {
		fmt.Println("error occurred : ", err, n)
		return ""
	}

	return string(data)
}
func main() {

	var topFive [5]int
	m := make(map[string]int)

	fmt.Println("2022년 조바이든 대통령취임 연설문")
	fmt.Println("특정 단어가 몇번 반복될까요?")

	sample := file_read("sample.txt")
	s := strings.Split(sample, " ")
	
	for _, v := range s {
		m[strings.Trim(v,"\n")]++
		if len(v) > 4{		// 4글자 이하는 for,in,the 같은 조사가 너무 많음
			checkTop(&topFive, m[v])
		}
		
	}

	fmt.Println("가장 많이 등장하는 단어 상위 5")
	for i := 0; i < 5; i++ {
		for k := range m {
			if topFive[i] == m[k] {
				fmt.Printf("[%s] : %d\n", k, m[k])
			}
		}
	}
}
func checkTop(lst *[5]int, num int) bool {
	for i, val := range lst {
		if num > val {
			lst[i] = num
			return true
		}
	}
	return false
}
