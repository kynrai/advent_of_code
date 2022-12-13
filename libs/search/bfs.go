package search

type Point struct {
	X, Y int
}

type BFS[T any] struct {
	grid        [][]T
	queue       []Point
	NeighbourFn func(T) bool
}

func NewBFS[T any]() *BFS[T] {
	return &BFS[T]{}
}

func (b *BFS[T]) Find() {

}
