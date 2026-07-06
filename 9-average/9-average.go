package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
)

func main() {
	nlist := []int{}
	fmt.Println("평균값/ 최빈값/ 중앙값/ 기하평균 구하기 ")
	fmt.Println("리스트를 랜덤 생성합니다")
	fmt.Println("완료!")

	for i := 0; i < 20; i++ {
		nlist = append(nlist, rand.Intn(20)+1)
	}
	fmt.Println("생성 결과 : ", nlist)
	slices.Sort(nlist)
	fmt.Println("정렬 : ", nlist)

	fmt.Println("평균값 : ", average(nlist))
	fmt.Println("최빈값 : ", mean(nlist))
	fmt.Println("중앙값 : ", median(nlist))
	fmt.Printf("기하평균 : %.2f\n", geo_mean(nlist))
}

// 평균값
func average(n []int) int {
	var total int
	size := len(n)
	for i := 0; i < size; i++ {
		total += i
	}
	return total / size
}

// 최빈값
func mean(n []int) []int {
	m := map[int]int{}
	result := []int{}
	var maxcount int

	for _, v := range n {
		m[v]++
		if m[v] > maxcount {
			maxcount = m[v]
		}
	}
	for k, _ := range m {
		if maxcount == m[k] {
			result = append(result, k)
		}
	}
	return result
}

// 중앙값. 홀수면 그값
func median(n []int) int {
	size := len(n)
	if n == nil || size == 0 {
		return -1
	}

	if size <= 2 {
		if size == 1 {
			return n[0]
		} else {
			return (n[0] + n[1]) / 2
		}
	}
	msize := size / 2
	if size%2 == 0 {
		return (n[msize-1] + n[msize]) / 2
	} else {
		return n[msize]
	}
}

// 기하평균
func geo_mean(n []int) float64 {
	size := len(n)
	var total float64 = 1.0
	for _, v := range n {
		total *= float64(v)
	}
	return math.Pow(total,1.0/float64(size)) // n0.n1.n2.....nx의 곱을 n제곱근으로 나눈것
}
