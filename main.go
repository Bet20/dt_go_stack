package main

type stack struct {
	nodes []interface{}
	size  int
}

func (s *stack) Push(value interface{}) {
	if len(s.nodes) < s.size {
		s.nodes = append(s.nodes, value)
	}
}

func (s stack) Peek() interface{} {
	return s.nodes[len(s.nodes)-1]
}

func (s *stack) Pop() interface{} {

	if s.IsEmpty() {
		return nil
	}

	ret := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]
	return ret
}

func (s *stack) IsEmpty() bool {
	return len(s.nodes) < 1
}

func (s *stack) IsFull() bool {
	return s.size == len(s.nodes)
}
