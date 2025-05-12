package queue

type Queue[T any] interface {
	Enqueue(data T)
	Dequeue() (T, error)
	IsEmpty() bool
	Length() int
	Peek() (T, bool)
	Clear()
}

type node[T any] struct {
	data T
	next *node[T]
}
