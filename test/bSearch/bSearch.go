// 알파벳 순서로 정렬된 단어리스트에 원하는 단어가 포함되어있는지
// 확인하는 이진탐색알고리즘을 작성해보세요

package main

import "fmt"

func main() {
	a := []string{"as", "at", "bbx", "crop", "diana", "eland", "golang", "kknd", "yong", "xls", "zztop"}
	var find string = "zztop"

	if result := bSearch(a, find); result == -1 {
		fmt.Println("can't find : ", find)
	} else {
		fmt.Printf("find! it's %dth element : %s\n", result, find)
	}

}
func bSearch(a []string, find string) int {
	first := 0
	last := len(a) - 1

	for last >= first {
		mid := (last + first) / 2 // save mid
		isSame := strCmp(a[mid], find)

		switch isSame {
		case 0:
			return mid // find!
		case -1:
			last = mid - 1 // a[mid] big, find smaller
		case 1:
			first = mid + 1 // a[mid] small, find bigger
		}
	}
	return -1 // not find
}
func strCmp(a string, b string) int {
	for i, _ := range a {
		if a[i] > b[i] {
			return -1
		} else if a[i] < b[i] {
			return 1
		}
	}
	return 0
}
