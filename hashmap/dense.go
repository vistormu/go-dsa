package hashmap

// store values in a dense slice while providing O(1) lookup by key.
//
// this type is not safe for concurrent use.
type DenseMap[K comparable, V any] struct {
	dense    []K
	values   []V
	location map[K]int
}

// create an empty densemap.
//
// time: O(1)
func NewDenseMap[K comparable, V any]() DenseMap[K, V] {
	return DenseMap[K, V]{
		dense:    make([]K, 0),
		values:   make([]V, 0),
		location: make(map[K]int),
	}
}

// add stores v under k.
//
// overwrite the existing value if k is already present.
//
// time: O(1) amortised
func (m *DenseMap[K, V]) Add(k K, v V) {
	row, ok := m.location[k]
	if ok {
		m.values[row] = v
		return
	}

	m.location[k] = len(m.dense)
	m.dense = append(m.dense, k)
	m.values = append(m.values, v)
}

// remove deletes k.
//
// return false if k does not exist.
//
// time: O(1)
func (m *DenseMap[K, V]) Remove(k K) bool {
	row, ok := m.location[k]
	if !ok {
		return false
	}

	last := len(m.dense) - 1
	lastK := m.dense[last]

	m.dense[row] = m.dense[last]
	m.values[row] = m.values[last]

	m.dense = m.dense[:last]
	m.values = m.values[:last]

	delete(m.location, k)
	if row != last {
		m.location[lastK] = row
	}

	return true
}

// get retrieves the stored value.
//
// return false if k does not exist.
//
// time: O(1)
func (m *DenseMap[K, V]) Get(k K) (V, bool) {
	row, ok := m.location[k]
	if !ok {
		var zero V
		return zero, false
	}
	return m.values[row], true
}

// getref retrieves a pointer to the stored value.
//
// return false if k does not exist.
//
// the returned pointer becomes invalid if the store grows and reallocates.
//
// time: O(1)
func (m *DenseMap[K, V]) GetRef(k K) (*V, bool) {
	row, ok := m.location[k]
	if !ok {
		return nil, false
	}
	return &m.values[row], true
}

// report whether k exists.
//
// time: O(1)
func (m *DenseMap[K, V]) Has(k K) bool {
	_, ok := m.location[k]
	return ok
}

// return the number of stored keys.
//
// time: O(1)
func (m *DenseMap[K, V]) Len() int {
	return len(m.dense)
}

// report whether no keys are stored.
//
// time: O(1)
func (m *DenseMap[K, V]) Empty() bool {
	return len(m.dense) == 0
}
