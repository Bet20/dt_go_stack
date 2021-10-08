package dt_go_stack

import "fmt"

func gorout(s *Stack) {

	s.Push("Hello")
	s.Push("Goodbye")
	s.Push("!")

	s.Pop()
	fmt.Printf("%s", s.Peek())

}
