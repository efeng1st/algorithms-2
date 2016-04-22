package sort

import (
	"math/rand"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestSort(t *testing.T) {
	data := ints
	a := IntSlice(data[0:])
	Sort(a, SHELL)
	if !IsSorted(a) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v ", data)
	}
}

func TestSortLarge(t *testing.T) {
	n := 1000
	intss := make([]int, n)
	for i := 0; i < len(intss); i++ {
		intss[i] = rand.Intn(1000)
	}
	data := intss
	a := IntSlice(data[0:])
	Sort(a, SHELL)
	if !IsSorted(a) {
		t.Errorf("sorted %v", intss)
		t.Errorf("   got %v ", data)
	}
}
