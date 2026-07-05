// 한국식 나이 + 만나이 계산식
// 한국식 나이 = 태어날때 1살 + 연도차이
// 만 나이 = 기본적으로 연도차이. 추가로 생일지났는지 여부에 따라 추가로 -1
package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type day struct {
	year  int
	month int
	day   int
}

func main() {
	var input string
	var birthday day
	var nowday day

	var y,m,d int

	fmt.Println("나이계산기")
	fmt.Print("생년월일 (yyyy-mm-dd) : ")
	fmt.Scanln(&input)

	// 2000-10-20 -> 20001020 -> int type 20001020
	sBirth := strings.Split(input, "-")

	var err error
	y, err = strconv.Atoi(sBirth[0])
	if err != nil {
		fmt.Println("can't convert year")
		return
	}
	m, err = strconv.Atoi(sBirth[1])
	if err != nil {
		fmt.Println("can't convert month")
		return
	}
	d, err = strconv.Atoi(sBirth[2])
	if err != nil {
		fmt.Println("can't convert day")
		return
	}

	birthday.year = y
	birthday.month = m
	birthday.day = d

	now := time.Now()
	nowday = day{
		year:  now.Year(),
		month: int(now.Month()),
		day:   now.Day(),
	}
	fmt.Println("한국식 나이 : ", nowday.year - birthday.year + 1)
	fmt.Println("만 나이 : ", man_nai_calc(birthday,nowday))
}
// 만 나이 계산은 총 2가지만 생각. 태어난 연도와 생일. 이번 연도에서 태어난 연도를 빼고, 올해 생일이 지나지 않았다면 추가로 1살을 더 빼서 만 나이를 쉽게 계산
func man_nai_calc(birth day, now day) int {
	var year = now.year - birth.year
	var other int = 0
	
	if birth.month > now.month{
		other = -1		
	}else if birth.month == now.month{
		if birth.day > now.day{
			other = -1
		}else{
			other = 0
		}
	}else{
		other = 0
	}
	return year + other
}