package algo

// compare to http://play.golang.org/p/Ma2GXvj3XP

func sort(a []int) {
	aux := make([]int, len(a)) // alloc it once, no allocs in recursive calls
	_sort(a, aux, 0, len(a)-1)
}

func _sort(a, aux []int, lo, hi int) {
	// recursive exit
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2 // start with lo+
	_sort(a, aux, lo, mid)
	_sort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a, aux []int, lo, mid, hi int) {
	// assert a[lo:mid] is sorted
	// assert a[mid+1:hi] is sorted

	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}
