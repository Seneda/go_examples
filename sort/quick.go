package sort

func Quick(l []int) []int {
	if len(l) <= 1 {
		return l
	}
 	less, more := Partition(l)
	return append(Quick(less), Quick(more)...)
}

func Partition(l []int) ([]int, []int){
	pivotindex := len(l)-1 // Choose the last element
	pivot := l[pivotindex]
	less := []int{}
	more := []int{pivot}
	for _, v := range append(l[:pivotindex], l[pivotindex+1:]...) {
		if v <= pivot {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}
	return less, more
}
