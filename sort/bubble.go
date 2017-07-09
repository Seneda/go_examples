package sort

func Bubble(l []int) []int {
    s := make([]int, len(l))
    copy(s, l)
    for i := 0; i < len(l); i++ {
        for j := 0; j < len(s) - 1; j++ {
            if s[j] > s[j+1] {
                s[j], s[j+1] = s[j+1], s[j]
            }
        }
    }
    return s
}