// bmi = 키/체중의 제곱
// 저체중: 18.5 미만정상: 18.5 ~ 22.9비만전단계(과체중): 23.0 ~ 24.91단계 비만: 25.0 ~ 29.92단계 비만: 30.0 ~ 34.93단계 비만: 35.0 이상
package main

import "fmt"

func main() {
	var height int
	var weight int

	fmt.Println("당신의 BMI를 체크합니다.")
	fmt.Println("키 (cm) : ")
	fmt.Scanf("%d\n", &height)
	fmt.Println("체중 (kg) : ")
	fmt.Scanf("%d\n", &weight)

	s := bmi_calc(height, weight)
	if s != "" {
		fmt.Println("error occured")
	}
}
func bmi_calc(height int, weight int) string {
	if height < 0 || weight < 0 {
		fmt.Println("wrong number under 0")
		return ""
	}
	bmi := float64(height / weight * weight)
	if bmi < 18.5 {
		return fmt.Sprintf("%.2f. 저체중입니다.", bmi)
	} else if bmi < 22.9 {
		return fmt.Sprintf("%.2f. 과체중입니다.", bmi)
	} else if bmi < 24.9 {
		return fmt.Sprintf("%.2f. 1단계 비만입니다.", bmi)
	} else if bmi < 29.2 {
		return fmt.Sprintf("%.2f. 2단계 비만입니다.", bmi)
	} else {
		return fmt.Sprintf("%.2f. 비만입니다", bmi)
	}
}
