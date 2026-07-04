// 섭씨-화씨 변환 프로그램
//섭씨(°C)를 화씨(°F)로 변환하는 공식은 F = (C * 1.8) + 32) , 화씨에서 섭씨로 변환하는 공식은 C = (F - 32) ÷ 1.8
package main

import "fmt"

func cel_feh(cel float64) float64 {
	return (cel * 1.8) + 32
}
func feh_cel(feh float64) float64 {
	return (feh - 32) / 1.8
}
func main() {
	var temperature float64
	var choice uint8

	fmt.Println("온도 변환 (섭씨-화씨)")
	fmt.Print("선택 (1:섭씨 -> 화씨 , 2:화씨 -> 섭씨) : ")
	fmt.Scanf("%d", &choice)
	fmt.Print("온도 : ")
	fmt.Scanf("%f\n", &temperature)
	

	if choice == 1 {
		t := cel_feh(temperature)
		fmt.Printf("섭씨 %.2f도 --> 화씨 %.2f도\n", temperature, t)
	} else if choice == 2 {
		t := feh_cel(temperature)
		fmt.Printf("화씨 %.2f도 --> 섭씨 %.2f도\n", temperature, t)
	} else {
		fmt.Println("입력 오류")
	}
}
