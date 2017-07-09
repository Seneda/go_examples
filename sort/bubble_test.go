package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	s := []int{1,2,3,5,7}
	l := Bubble([]int{2,3,1,5,7})

	for i, v := range s {
		if v != l[i] {
			t.Error(fmt.Sprintf("List is not sorted: %d", l))
		}
	}
}

func TestBubbleLong(t *testing.T) {
	s := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21}
	l := Bubble([]int{21,20,19,18,17,16,15,14,13,12,11,10,9,8,7,6,5,4,3,2,1})

	for i, v := range s {
		if v != l[i] {
			t.Error(fmt.Sprintf("List is not sorted: %d", l))
		}
	}
}
