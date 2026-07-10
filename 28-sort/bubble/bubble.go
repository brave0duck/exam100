package bubble

// [5,3,3,2,1] => [1,2,3,3,5]
func SortBubble(n *[]int) []int{

	size := len(n)

	for i:=0 ; i<size ; i++ {
		for j:=i; j<size ; j++{
			if n[j] > n[j+1]{
				n[j],n[j+1] = n[j+1],n[j]
			}
		}
	}
	return n
}	