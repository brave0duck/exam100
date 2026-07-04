// bmi계산 프로그램
// 저체중: 18.5 미만
// 정상: 18.5 ~ 22.9
// 비만전단계(과체중): 23.0 ~ 24.91단계
// 비만: 25.0 ~ 29.92단계
// 비만: 30.0 ~ 34.93단계
// 비만: 35.0 이상
package main

import "fmt"

func main() {
	var height int
	var weight int

	fmt.Println("당신의 BMI를 체크합니다.")
	fmt.Print("키 (cm) : ")
	fmt.Scanf("%d\n", &height)
	fmt.Print("체중 (kg) : ")
	fmt.Scanf("%d\n", &weight)

	s := bmi_calc(height, weight)
	if s == "" {
		fmt.Println("error occured")
	} else {
		fmt.Println(s)
	}
}
func double_weight(n int) float64 {
	w := float64(n) / 100
	return w * w
}
func bmi_calc(height int, weight int) string {
	if height < 0 || weight < 0 {
		fmt.Println("wrong number under 0")
		return ""
	}
	var message string
	bmi := float64(weight) / double_weight(height)

	if bmi < 18.5 {
		message = "[ 저체중 ] 입니다."
	} else if bmi < 22.9 {
		message = "[ 정상 ] 입니다."
	} else if bmi < 25 {
		message = "[ 과체중 ] 입니다."
	} else if bmi < 29.2 {
		message = "[ 1단계 비만 ] 입니다."
	} else if bmi < 34.9 {
		message = "[ 2단계 비만 ] 입니다."
	} else {
		message = "[ 3단계 비만 ] 입니다."
	}
	return fmt.Sprintf("당신의 BMI는 %.2f이며 %s 정상체중은 %.2f - %.2f kg 입니다",
		bmi, message, 18.5*double_weight(height), 22.9*double_weight(height))
}
