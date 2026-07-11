package main

import (
	"28-sort/Bubble"
	"28-sort/Selection"
	"fmt"
	"math/rand/v2"
)

func main() {
	var input int8
	sample := make([]int, 0, 20)

	for range 20 {
		sample = append(sample, rand.IntN(100)+1)
	}
	fmt.Println(sample)

	fmt.Println("---sorting---")
	fmt.Println("1. bubble sort")
	fmt.Println("2. select sort")
	fmt.Println("3. insert sort")
	fmt.Println("4. binary sort")
	fmt.Println("0. exit")
	fmt.Println(">>")
	fmt.Scanf("%d\n", &input)

	switch input {
	case 1:
		{
			Bubble.Sort(&sample)
			fmt.Println(sample)
		}
	case 2:
		{
			Selection.Sort(&sample)
			fmt.Println(sample)
		}
	default:
		fmt.Println("wrong input")
	}
}
