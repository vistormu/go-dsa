package queue

type Queue[T any] interface {
	Enqueue(data T)
	Dequeue() (T, error)
	Peek() (T, bool)
	ToSlice() []T
	IsEmpty() bool
	Length() int
	Clear()
}

type node[T any] struct {
	data T
	next *node[T]
}
