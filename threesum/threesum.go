package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/hebl/algorithms/stdlib"
)

// ThreeSum 三元组求和为0问题，暴力求解
func ThreeSum(a []int) int {
	defer stdlib.TimeTrack(time.Now(), "ThreeSum")
	N := len(a)
	count := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				if (a[i] + a[j] + a[k]) == 0 {
					count++
				}
			}
		}
	}

	return count
}

// ThreeSumFast 三元组求和为0问题，快速解法
func ThreeSumFast(a []int) int {
	defer stdlib.TimeTrack(time.Now(), "ThreeSumFast")
	sort.Ints(a)

	N := len(a)
	count := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if stdlib.BinarySearch(-a[i]-a[j], a) > j {
				count++
			}
		}
	}
	return count
}

func main() {
	a, _ := stdlib.ReadAllInts()

	fmt.Println(ThreeSum(a))
	fmt.Println(ThreeSumFast(a))
}
