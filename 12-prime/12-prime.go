// 소수판별기
/*
특정 숫자 n이 소수인지 판별하는 가장 효율적인 로직은
2부터 sqrt{n}까지의 정수로 n을 나누어 보는 것입니다.
sqrt{n}까지만 확인해도 약수의 대칭성에 의해 소수 여부를 정확히 판단할 수 있어 불필요한 연산을 대폭 줄일 수 있습니다.

소수 판별 로직 (제곱근 활용):

1. 1이하의 숫자는 소수가 아니므로 바로 False를 반환합니다.
2. 반복문 수행: \(2\)부터 \(\sqrt{N}\)까지의 수 \(i\)로 \(N\)을 나눕니다.
3. 나머지 확인: \(N \% i == 0\)인 경우(나누어 떨어지는 경우) 약수가 존재하므로 소수가 아닙니다(False).
4. 소수 판정: 반복문이 끝날 때까지 나누어 떨어지는 수가 없다면 소수입니다(True).
*/
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	t := int(math.Sqrt(float64(n)))

	var prime bool = true

	for i := 2; i <= t; i++ {
		if n%i == 0 {
			return false
		}
	}
	return prime
}
func main() {
	var num int
	fmt.Println("prime number checker")
	fmt.Print("input number : ")
	fmt.Scanf("%d", &num)

	if isPrime(num) {
		fmt.Println(num, "is prime")
	} else {
		fmt.Println(num, "is not prime")
	}
}
