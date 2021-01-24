package runtime

// Stack of integers acting as memory
type Stack struct {
	data []byte
}

func (s *Stack) Push(i byte) {
	s.data = append(s.data, i)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		panic("Stack empty")
	}

	lastIndex := len(s.data) - 1
	ret := s.data[lastIndex]

	s.data = s.data[:lastIndex]

	return ret
}
