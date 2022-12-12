package queue

type Queue[T any] []T

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Push(str T) {
	*q = append(*q, str)
}

func (q *Queue[T]) Pop() (T, bool) {
	var n T
	if q.IsEmpty() {
		return n, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}
