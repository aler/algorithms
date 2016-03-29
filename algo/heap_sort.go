package algo

import (
	"log"
)

func heapSort(a []rune) {
	// make array heap ordered
	for k := len(a) / 2; k >= 1; k-- {
		sink(a, k)
	}

	assertHeap(a)

	// sort
	for len(a) > 1 {
		exch(a, 1, len(a))
		a = a[:len(a)-1]
		sink(a, 1)
	}
}

func sink(a []rune, i int) {
	for i*2 <= len(a) {
		c := i * 2
		max := c
		if c < len(a) && less(a, c, c+1) {
			max = c + 1
		}
		if !less(a, i, max) {
			break
		}
		exch(a, i, max)
		i = max
	}
}

func less(a []rune, i, j int) bool {
	// decrement indicies to works with regular array
	// heap array starts with index 1 to simplify math
	if a[i-1] < a[j-1] {
		return true
	}
	return false
}

func exch(a []rune, i, j int) {
	// decrement indicies to works with regular array
	// heap array starts with index 1 to simplify math
	a[i-1], a[j-1] = a[j-1], a[i-1]
}

func assertHeap(a []rune) {
	i := len(a)
	for i > 1 {
		if !less(a, i, i/2) {
			log.Fatalf("\nisHeap: %d, %q, %q\n", i, a[i], a[i/2])
		}
		i--
	}
}
