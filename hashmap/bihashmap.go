package hashmap

type BiHashmap[K, V comparable] struct {
	forward  map[K]V
	backward map[V]K
}

func NewBiHashmap[K, V comparable]() *BiHashmap[K, V] {
	return &BiHashmap[K, V]{
		forward:  make(map[K]V),
		backward: make(map[V]K),
	}
}

func (b *BiHashmap[K, V]) Put(k K, v V) {
	// if the key already exists, remove the old value
	oldValue, ok := b.forward[k]
	if ok {
		delete(b.backward, oldValue)
	}

	// if the value already exists, remove the old key
	oldKey, ok := b.backward[v]
	if ok {
		delete(b.forward, oldKey)
	}

	// add the new key-value pair
	b.forward[k] = v
	b.backward[v] = k
}

func (b *BiHashmap[K, V]) GetByKey(k K) (V, bool) {
	v, ok := b.forward[k]
	return v, ok
}

func (b *BiHashmap[K, V]) GetByValue(v V) (K, bool) {
	k, ok := b.backward[v]
	return k, ok
}

func (b *BiHashmap[K, V]) RemoveByKey(k K) {
	v, ok := b.forward[k]
	if ok {
		delete(b.backward, v)
		delete(b.forward, k)
	}
}

func (b *BiHashmap[K, V]) RemoveByValue(v V) {
	k, ok := b.backward[v]
	if ok {
		delete(b.forward, k)
		delete(b.backward, v)
	}
}

func (b *BiHashmap[K, V]) Keys() []K {
	keys := make([]K, 0, len(b.forward))
	for k := range b.forward {
		keys = append(keys, k)
	}
	return keys
}

func (b *BiHashmap[K, V]) Values() []V {
	values := make([]V, 0, len(b.backward))
	for v := range b.backward {
		values = append(values, v)
	}
	return values
}

func (b *BiHashmap[K, V]) Length() int {
	return len(b.forward)
}

func (b *BiHashmap[K, V]) Clear() {
	b.forward = make(map[K]V)
	b.backward = make(map[V]K)
}
