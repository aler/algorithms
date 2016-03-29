package algo

import "math/rand"

func shuffle(a []rune) {
	// http://stackoverflow.com/questions/12264789/shuffle-array-in-go/
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

func quickSort(a []rune) {
	shuffle(a)
	_quickSort(a)
}

func _quickSort(a []rune) {
	if len(a) <= 1 {
		return
	}

	j := partition(a)
	_quickSort(a[:j])
	_quickSort(a[j+1:])
}

func partition(a []rune) int {
	v, i, j := a[0], 1, len(a)-1
	for true {
		for a[i] < v && i != len(a)-1 {
			i++
		}
		for a[j] > v && j != 0 {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[0], a[j] = a[j], a[0]
	return j
}
