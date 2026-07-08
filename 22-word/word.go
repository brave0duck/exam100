// 가장 많이 등장한 단어 찾기
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

	var topTier [5]int
	m := make(map[string]int)

	fmt.Println("2022년 조바이든 대통령취임 연설문")
	fmt.Println("특정 단어가 몇번 반복될까요?")

	sample := file_read("sample.txt")
	s := strings.Split(sample, " ")

	for _, v := range s {
		m[v]++
		checkTop(&topTier, m[v])
	}

	fmt.Println("가장 많이 등장하는 단어 상위 5")
	for i := 0; i < 5; i++ {
		for k, _ := range m {
			if topTier[i] == m[k] {
				fmt.Printf("[%s] : %d\n", k, m[k])
			}
		}
	}
}

func checkTop(t *[5]int, num int) bool {
	for k, v := range t {
		if num > v {
			t[k] = num
			return true
		}
	}
	return false
}
