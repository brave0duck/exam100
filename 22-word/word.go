// 가장 많이 등장한 단어 찾기
package main

import (
	"fmt"
	"os"
	"strings"
	"slices"
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
	var max_word string
	max_count := make([]int, 5)
	m := make(map[string]int, 0)

	fmt.Println("2022년 조바이든 대통령취임 연설문")
	fmt.Println("특정 단어가 몇번 반복될까요?")

	sample := file_read("sample.txt")
	s := strings.Split(sample, " ")

	for _, v := range s {
		m[v]++
		if slices.Contains(max_count,m[v]){
			
		} > max_count {
			max_count = m[v]
			max_word = v
		}
	}
	fmt.Printf("가장 많이 등장하는 단어는 %s이고 %d회 등장했습니다\n", max_word, max_count)
}
func is
