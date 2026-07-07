// 문자열 빈도 프로그램 - 파일을 열어서 특정 문자열이 몇번 반복되는지 확인
// 1. 파일 열기- 읽기 - 카운트
package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("2022년 조바이든 대통령취임 연설문")
	fmt.Println("특정 단어가 몇번 반복될까요?")
	fmt.Print("찾을 단어 : ")

	if scanner.Scan() {
		input := scanner.Text()

		sample := file_read("sample.txt")

		n := strings.Count(sample, input)
		if n > 0 {
			fmt.Printf("==> %d개 있습니다.\n", n)
		} else {
			fmt.Println("==> 단어가 없습니다. X")
		}
	}
}
