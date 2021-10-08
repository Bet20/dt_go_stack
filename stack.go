package dt_go_stack

type Stack struct {
	Nodes []interface{}
	Size  int
}

// Pushes value of interface to the top of the stack
func (s *Stack) Push(value interface{}) {
	if len(s.Nodes) < s.Size {
		s.Nodes = append(s.Nodes, value)
	}
}

// Returns the value of the top of the stack with out poping it
func (s Stack) Peek() interface{} {
	return s.Nodes[len(s.Nodes)-1]
}

// Pops the value in top of the stack returning it
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ret := s.Nodes[len(s.Nodes)-1]
	s.Nodes = s.Nodes[:len(s.Nodes)-1]
	return ret
}

// Checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.Nodes) < 1
}

// Checks if the stack is full
func (s *Stack) IsFull() bool {
	return s.Size == len(s.Nodes)
}
