// O(n2). selection sort 
// 1부터 n까지 최저값을 찾고 맨처음으로 스왑
// 2부터 n까지 최저값을 찾고 두번째와 스왑 ..
// <bubble sort와의 차이점?>
// 안쪽 for문에서 스왑은 1번으로 고정( 스왑횟수가 적음)

package Selection

func Sort(n *[]int) {
	var isChange bool
	size := len(*n)
	
	for i:=0; i<size; i++{
		min := i
		for j := i+1; j <size; j++{
			if (*n)[min] > (*n)[j]{
				min = j
				isChange = true
			}
		}
		if isChange {
			(*n)[i] ,(*n)[min] = (*n)[min], (*n)[i]
		}
		isChange = false
	}

}