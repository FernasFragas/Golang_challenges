package stacks

type Stack struct {
	items []byte
}

func (s *Stack) Push(value byte) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() byte {
	if len(s.items) == 0 {
		return byte(0)
	}
	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return value
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
