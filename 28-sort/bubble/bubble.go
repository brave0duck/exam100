// O(n2).bubble sort
// 1부터 n까지 왼쪽부터 오른쪽으로 비교하면서 작은값을 왼쪽으로 이동
// 2부터 n까지 ...
// 3부터 n까지 ...
// 비교와 스왑연산의 동시수행으로 효율이 떨어지는 알고리즘
package Bubble

func Sort(n *[]int) {
	size := len(*n)

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			if (*n)[j] > (*n)[j+1] {
				(*n)[j], (*n)[j+1] = (*n)[j+1], (*n)[j]
			}
		}
	}
}
