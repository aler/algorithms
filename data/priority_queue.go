package data

type priorityQueue struct {
	ar []int
}

func newPriorityQueue() *priorityQueue {
	return &priorityQueue{ar: make([]int, 1)}
}

func (q *priorityQueue) insert(val int) {
	q.ar = append(q.ar, val)
	q.swim(len(q.ar) - 1)
}

func (q *priorityQueue) max() int {
	return q.ar[1]
}

func (q *priorityQueue) deleteMax() int {
	max := q.ar[1]
	q.ar[1] = q.ar[len(q.ar)-1]
	q.ar = q.ar[:len(q.ar)-1]
	q.sink(1)
	return max
}

func (q *priorityQueue) swim(i int) {
	for i > 1 && q.ar[i/2] < q.ar[i] {
		q.ar[i/2], q.ar[i] = q.ar[i], q.ar[i/2]
		i /= 2
	}
}

func (q *priorityQueue) sink(i int) {
	c := i * 2
	for c <= len(q.ar)-1 {
		// find max idx
		max := c
		if c < len(q.ar)-1 && q.ar[c] < q.ar[c+1] {
			max = c + 1
		}
		// if parent greater then children max, done
		if q.ar[i] > q.ar[max] {
			break
		}
		// exchange and update child idx
		q.ar[i], q.ar[max] = q.ar[max], q.ar[i]
		c = max
	}
}
