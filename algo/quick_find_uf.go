package algo

type quickFindUF struct {
	id    []int
	comps int
}

func newQuickFindUF(n int) *quickFindUF {
    id := make([]int, n)
    for i := range id {
        id[i] = i
    }
    return &quickFindUF{id: id, comps: n} 
}

func (s *quickFindUF) union(p, q int) {
    pid, qid := s.find(p), s.find(q)
    if pid == qid {
        return
    }
    for _, v := range s.id {
        if v == pid {
            v = qid
        }
    }
    s.comps--
}

func (s *quickFindUF) find(p int) int {
    return s.id[p]
}

func (s *quickFindUF) connected(p, q int) bool {
    return s.find(p) == s.find(q)
}

func (s *quickFindUF) count() int {
    return s.comps
}
