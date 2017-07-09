package sort

func Bubble(l []int) []int {
    for i := 0; i < len(l); i++ {
        for j := 0; j < len(l) - 1; j++ {
            if l[j] > l[j+1] {
                l[j], l[j+1] = l[j+1], l[j]
            }
        }
    }
    return l
}