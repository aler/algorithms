package algo

type quickUnionUFWeighted struct {
	id    []int
	sz    []int
	comps int
}

func newQuickUnionWeightedUF(n int) *quickUnionUFWeighted {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}
	return &quickUnionUFWeighted{id: id, sz: sz, comps: n}
}

func (s *quickUnionUFWeighted) union(p, q int) {
	i, j := s.find(p), s.find(q)
	if i == j {
		return
	}
	if s.sz[i] < s.sz[j] {
		s.id[i] = j
		s.sz[j] += s.sz[i]
	} else {
		s.id[j] = i
		s.sz[i] += s.sz[j]
	}
	s.comps--
}

func (s *quickUnionUFWeighted) find(p int) int {
	for p != s.id[p] {
		p = s.id[p]
	}
	return p
}

func (s *quickUnionUFWeighted) connected(p, q int) bool {
	return s.find(p) == s.find(q)
}

func (s *quickUnionUFWeighted) count() int {
	return s.comps
}
