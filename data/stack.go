package data

import "log"

type stack struct {
	first *list
}

func (s *stack) push(item int) {
	s.first = &list{item: item, next: s.first}
}

func (s *stack) pop() int {
	if s.first == nil {
		log.Fatal("stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	return item
}

type resizingArrayStack struct {
    items []int
}

func newResizingArrayStack() *resizingArrayStack {
    return &resizingArrayStack{items: make([]int, 1)}
} 

func (s *resizingArrayStack) push(item int) {
    s.items = append(s.items, item)
}

func (s *resizingArrayStack) pop() int {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]

    // shrink if needed    
    if len(s.items) == cap(s.items)/4 {
        items := make([]int, len(s.items), cap(s.items)/2)
        copy(items, s.items)
        s.items = items
    }
    return item
}
