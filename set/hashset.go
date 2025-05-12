package set

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: make(map[T]struct{})}
}

func (h *HashSet[T]) Add(value T) {
	h.data[value] = struct{}{}
}

func (h *HashSet[T]) Remove(value T) {
	delete(h.data, value)
}

func (h *HashSet[T]) Contains(value T) bool {
	_, exists := h.data[value]
	return exists
}

func (h *HashSet[T]) Clear() {
	h.data = make(map[T]struct{})
}

func (h *HashSet[T]) Size() int {
	return len(h.data)
}

func (h *HashSet[T]) Values() []T {
	values := make([]T, 0, len(h.data))
	for key := range h.data {
		values = append(values, key)
	}
	return values
}

func (h *HashSet[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *HashSet[T]) Union(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()
	for key := range h.data {
		result.Add(key)
	}
	for key := range other.data {
		result.Add(key)
	}
	return result
}

func (h *HashSet[T]) Intersection(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()
	for key := range h.data {
		if other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}

func (h *HashSet[T]) Difference(other *HashSet[T]) *HashSet[T] {
	result := NewHashSet[T]()
	for key := range h.data {
		if !other.Contains(key) {
			result.Add(key)
		}
	}
	return result
}
