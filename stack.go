package dt_go_stack

type Stack struct {
	Nodes []interface{}
	Size  int
}

func (s *Stack) Push(value interface{}) {
	if len(s.Nodes) < s.Size {
		s.Nodes = append(s.Nodes, value)
	}
}

func (s Stack) Peek() interface{} {
	return s.Nodes[len(s.Nodes)-1]
}

func (s *Stack) Pop() interface{} {

	if s.IsEmpty() {
		return nil
	}

	ret := s.Nodes[len(s.Nodes)-1]
	s.Nodes = s.Nodes[:len(s.Nodes)-1]
	return ret
}

func (s *Stack) IsEmpty() bool {
	return len(s.Nodes) < 1
}

func (s *Stack) IsFull() bool {
	return s.Size == len(s.Nodes)
}
