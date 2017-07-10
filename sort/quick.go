package sort

func Quick(l []int) []int {
	if len(l) <= 1 {
		return l
	}
 	less, more := Partition(l)
	return append(Quick(less), Quick(more)...)
}

func Partition(l []int) ([]int, []int){
	pivot := l[len(l)-1]
	less := []int{}
	more := []int{pivot}
	for _, v := range l[:len(l)-1] {
		if v <= pivot {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}
	return less, more
}
