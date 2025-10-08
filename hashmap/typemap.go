package hashmap

import (
	"iter"
	"reflect"
	"sync"
)

// `TypeMap` is a concurrent map keyed by type (reflect.Type)
//
// values are always stored as pointers to the data to allow in-place mutation
//
// there can only exist one value per type
type TypeMap struct {
	mu   sync.RWMutex
	data map[reflect.Type]any
}

// `NewTypeMap` creates and returns an empty, concurrency-safe `TypeMap`
func NewTypeMap() *TypeMap {
	return &TypeMap{
		data: make(map[reflect.Type]any),
	}
}

// `Len` returns the current number of entries in the map
//
// safe for concurrent use
func (tm *TypeMap) Len() int {
	tm.mu.RLock()
	n := len(tm.data)
	tm.mu.RUnlock()

	return n
}

// `Keys` returns a snapshot of all base-type keys currently stored
//
// the returned slice is a copy and is safe to use without additional locking
func (tm *TypeMap) Keys() []reflect.Type {
	tm.mu.RLock()

	keys := make([]reflect.Type, 0, len(tm.data))

	for k := range tm.data {
		keys = append(keys, k)
	}
	tm.mu.RUnlock()

	return keys
}

// `Values` returns a snapshot of all stored values (each is a pointer to its base type)
//
// the returned slice is a copy and is safe to use without additional locking
func (tm *TypeMap) Values() []any {
	tm.mu.RLock()

	values := make([]any, 0, len(tm.data))

	for _, v := range tm.data {
		values = append(values, v)
	}
	tm.mu.RUnlock()

	return values
}

// `Iter` returns an iterator over a snapshot of the current contents
//
// it does NOT hold locks while yielding (to avoid deadlocks / long critical sections)
//
// each yielded value is the stored pointer `*BaseType`
//
// it does not guarantee the same order on each call
//
// example:
//
//	for k, v := range tm.Iter() {
//	    // k is reflect.Type (base type), v is pointer to k
//	}
func (tm *TypeMap) Iter() iter.Seq2[reflect.Type, any] {
	return func(yield func(rt reflect.Type, v any) bool) {
		type pair struct {
			k reflect.Type
			v any
		}

		tm.mu.RLock()
		pairs := make([]pair, 0, len(tm.data))
		for k, v := range tm.data {
			pairs = append(pairs, pair{k: k, v: v})
		}
		tm.mu.RUnlock()

		for _, p := range pairs {
			if !yield(p.k, p.v) {
				return
			}
		}
	}
}

// `Clear` removes all entries from the map
//
// safe for concurrent use
func (tm *TypeMap) Clear() {
	tm.mu.Lock()
	tm.data = make(map[reflect.Type]any)
	tm.mu.Unlock()
}

// =======
// helpers
// =======
func baseType(t reflect.Type) reflect.Type {
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

func asPointerToValue(v any) any {
	rv := reflect.ValueOf(v)
	rt := rv.Type()
	if rt.Kind() == reflect.Pointer {
		return v
	}
	ptr := reflect.New(rt)
	ptr.Elem().Set(rv)
	return ptr.Interface()
}

// =================
// generic functions
// =================

// `Add[T]` stores `v` under the base type of `T`
//
// if a type of `*T` is passed, it will  be stored internally as `reflect.TypeFor[T]`,
//
// the value is stored as a reference inside the map to allow in-place mutation
//
// safe for concurrent use
func Add[T any](tm *TypeMap, v T) {
	rt := reflect.TypeFor[T]()
	key := baseType(rt)
	ptr := asPointerToValue(v)

	tm.mu.Lock()
	tm.data[key] = ptr
	tm.mu.Unlock()
}

// `Get[T]` retrieves the pointer to the stored value for `T` (the base type)
//
// it returns `nil, false` if there is no data under the specified type or if the type cast to `*T` fails
//
// safe for concurrent use
func Get[T any](tm *TypeMap) (*T, bool) {
	rt := reflect.TypeFor[T]()
	key := baseType(rt)

	tm.mu.RLock()
	v, ok := tm.data[key]
	tm.mu.RUnlock()

	if !ok {
		return nil, false
	}

	out, ok := v.(*T)

	return out, ok
}

// `Remove[T]` removes the entry for the base type `T`
//
// it returns `false` if there is no value under the type `T`
//
// safe for concurrent use
func Remove[T any](tm *TypeMap) bool {
	rt := reflect.TypeFor[T]()
	key := baseType(rt)

	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, ok := tm.data[key]; !ok {
		return false
	}

	delete(tm.data, key)

	return true
}

// ====================
// reflection functions
// ====================

// `AddByType` stores `v` under the provided type `t`, where the key is the base type of `t`
//
// returns `false` is the types are incompatible
//
// safe for concurrent use
func AddByType(tm *TypeMap, v any, t reflect.Type) bool {
	if t == nil || v == nil {
		return false
	}

	key := baseType(t)
	if key == nil {
		return false
	}
	ptrType := reflect.PointerTo(key)

	rv := reflect.ValueOf(v)
	rt := rv.Type()

	tm.mu.Lock()
	defer tm.mu.Unlock()

	if rt == ptrType || rt.AssignableTo(ptrType) {
		if rt != ptrType && rv.CanConvert(ptrType) {
			rv = rv.Convert(ptrType)
		}
		tm.data[key] = rv.Interface()
		return true
	}

	if rt == key || rt.AssignableTo(key) || rv.CanConvert(key) {
		if rt != key {
			if !rv.CanConvert(key) {
				return false
			}
			rv = rv.Convert(key)
		}
		ptr := reflect.New(key)
		ptr.Elem().Set(rv)
		tm.data[key] = ptr.Interface()
		return true
	}

	return false
}

// `GetByType` retrieves the stored pointer for the provided type t's base type
//
// it returns (nil, false) if t is nil or no value exists
//
// safe for concurrent use
func GetByType(tm *TypeMap, t reflect.Type) (any, bool) {
	if t == nil {
		return nil, false
	}

	key := baseType(t)

	tm.mu.RLock()
	raw, ok := tm.data[key]
	tm.mu.RUnlock()

	return raw, ok
}

// `DeleteByType` removes the entry keyed by base type of t
//
// returns false if t is nil or no such entry exists
//
// safe for concurrent use
func DeleteByType(tm *TypeMap, t reflect.Type) bool {
	if t == nil {
		return false
	}

	key := baseType(t)

	tm.mu.Lock()
	_, ok := tm.data[key]
	if !ok {
		tm.mu.Unlock()
		return false
	}

	delete(tm.data, key)

	tm.mu.Unlock()

	return true
}
