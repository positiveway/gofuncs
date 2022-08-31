package gofuncs

import (
	"sync"
)

type Map[K comparable, V any] struct {
	mapping map[K]V
}

func MakeMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		mapping: map[K]V{},
	}
}

func (m *Map[K, V]) Put(key K, value V) {
	m.mapping[key] = value
}

func (m *Map[K, V]) CheckAndGet(key K) (V, bool) {
	value, present := m.mapping[key]
	return value, present
}

func (m *Map[K, V]) Get(key K) V {
	value, _ := m.CheckAndGet(key)
	return value
}

func (m *Map[K, V]) RangeOverShallowCopy(elementHandler func(key K, value V)) {
	copiedMap := ShallowCopyMap(m.mapping)

	for k, v := range copiedMap {
		elementHandler(k, v)
	}
}

func (m *Map[K, V]) RangeOverDeepCopy(elementHandler func(key K, value V)) {
	copiedMap := Copy(m.mapping)

	for k, v := range copiedMap {
		elementHandler(k, v)
	}
}

func (m *Map[K, V]) Pop(key K) V {
	value := m.mapping[key]
	delete(m.mapping, key)
	return value
}

type ThreadSafeMap[K comparable, V any] struct {
	mapping map[K]V
	mutex   sync.Mutex
}

func MakeThreadSafeMap[K comparable, V any]() *ThreadSafeMap[K, V] {
	return &ThreadSafeMap[K, V]{
		mapping: map[K]V{},
	}
}

func (threadMap *ThreadSafeMap[K, V]) Put(key K, value V) {
	threadMap.mutex.Lock()
	defer threadMap.mutex.Unlock()

	threadMap.mapping[key] = value
}

func (threadMap *ThreadSafeMap[K, V]) CheckAndGet(key K) (V, bool) {
	threadMap.mutex.Lock()
	defer threadMap.mutex.Unlock()

	value, present := threadMap.mapping[key]
	return value, present
}

func (threadMap *ThreadSafeMap[K, V]) Get(key K) V {
	value, _ := threadMap.CheckAndGet(key)
	return value
}

func (threadMap *ThreadSafeMap[K, V]) RangeOverShallowCopy(elementHandler func(key K, value V)) {
	threadMap.mutex.Lock()

	copiedMap := ShallowCopyMap(threadMap.mapping)

	threadMap.mutex.Unlock()

	for k, v := range copiedMap {
		elementHandler(k, v)
	}
}

func (threadMap *ThreadSafeMap[K, V]) RangeOverDeepCopy(elementHandler func(key K, value V)) {
	threadMap.mutex.Lock()

	copiedMap := Copy(threadMap.mapping)

	threadMap.mutex.Unlock()

	for k, v := range copiedMap {
		elementHandler(k, v)
	}
}

func (threadMap *ThreadSafeMap[K, V]) Pop(key K) V {
	threadMap.mutex.Lock()
	defer threadMap.mutex.Unlock()

	value := threadMap.mapping[key]
	delete(threadMap.mapping, key)
	return value
}
