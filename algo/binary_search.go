package algo

// BinarySearch searches sorted array a for a value v
// and returns index of that value if found with true ok flag
// if value is not found it returns false ok flag
func BinarySearch(a []int, v int) (int, bool) {
	lo, hi := 0, len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if v < a[mid] {
            hi = mid - 1
		} else if v > a[mid] {
            lo = mid + 1
		} else {
			return mid, true
		}
	}
	return -1, false
}
