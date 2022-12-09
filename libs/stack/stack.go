package stack

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(str T) {
	*s = append(*s, str)
}

func (s *Stack[T]) Pop() (T, bool) {
	var n T
	if s.IsEmpty() {
		return n, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
