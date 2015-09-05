package stack

import (
	"bytes"
	"strconv"
)

type Stack struct {
	values []int
	count  int
}

func NewStack() *Stack {
	return &Stack{values: make([]int, 8)}
}

func (s *Stack) Push(v int) {
	if s.count >= cap(s.values) {
		s.values = append(s.values, v)
	} else {
		s.values[s.count] = v
	}
	s.count++
}

func (s *Stack) Pop() (int, bool) {
	if s.count == 0 {
		return 0, true
	}
	s.count--
	return s.values[s.count], false
}

func (s *Stack) Count() int {
	return s.count
}

func (s *Stack) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i := 0; i < s.count; i++ {
		buf.WriteString(strconv.Itoa(s.values[i]))
		if i+1 < s.count {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]@")
	buf.WriteString(strconv.Itoa(s.count))
	return buf.String()
}
