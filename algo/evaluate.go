package algo

import "log"
import "strings"
import "strconv"

type (
	stringList struct {
		item string
		next *stringList
	}

	stringStack struct {
		first *stringList
	}

	intList struct {
		item int
		next *intList
	}

	intStack struct {
		first *intList
	}
)

func newStringStack() *stringStack {
	return &stringStack{first: nil}
}

func (s *stringStack) push(item string) {
	s.first = &stringList{item: item, next: s.first}
}

func (s *stringStack) pop() string {
	if s.first == nil {
		log.Fatal("stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	return item
}

func newIntStack() *intStack {
	return &intStack{first: nil}
}

func (s *intStack) push(item int) {
	s.first = &intList{item: item, next: s.first}
}

func (s *intStack) pop() int {
	if s.first == nil {
		log.Fatal("stack is empty")
	}
	item := s.first.item
	s.first = s.first.next
	return item
}

func evaluate(exp string) int {
	vals := newIntStack()
	ops := newStringStack()

	for _, c := range strings.Split(exp, " ") {
		switch c {
		case "(":
			break
		case "+", "-", "*", "/":
			ops.push(c)
		case ")":
			op := ops.pop()
			v1, v2 := vals.pop(), vals.pop()
			switch op {
			case "+":
				vals.push(v1 + v2)
			case "-":
				vals.push(v1 - v2)
			case "*":
				vals.push(v1 * v2)
			case "/":
				vals.push(v1 / v2)
			}
		default:
			val, _ := strconv.Atoi(c)
			vals.push(val)
		}
	}

	return vals.pop()
}
