package stdlib

//BinarySearch 整数的二分查找
// a 为一排序好的数组
func BinarySearch(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	var mid int

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
