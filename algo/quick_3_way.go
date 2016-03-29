package algo

func quick3WaySort(a []rune) {
    if len(a) <= 1 {
        return
    }
    lt, i, gt := 0, 1, len(a) - 1
    v := a[0]
    for i <= gt {
       if a[i] < v {
           q3exch(a, lt, i)
           lt++
           i++
       } else if a[i] > v {
           q3exch(a, gt, i)
           gt--
       } else {
           i++
       }
    }
    quick3WaySort(a[:lt])
    quick3WaySort(a[gt+1:])
}

func q3exch(a []rune, i, j int) {
	a[i], a[j] = a[j], a[i]
}