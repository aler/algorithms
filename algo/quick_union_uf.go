package algo

type quickUnionUF struct {
	id    []int
	comps int
}

func newQuickUnionUF(n int) *quickUnionUF {
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    return &quickUnionUF{id: id, comps: n} 
}

func (s *quickUnionUF) union(p, q int) {
    proot, qroot := s.find(p), s.find(q)
    if proot == qroot {
        return
    }
    s.id[proot] = qroot
    s.comps--
}

func (s *quickUnionUF) find(p int) int {
    for p != s.id[p] {
        p = s.id[p]
    }
    return p
}

func (s *quickUnionUF) connected(p, q int) bool {
    return s.find(p) == s.find(q)
}

func (s *quickUnionUF) count() int {
    return s.comps
}
