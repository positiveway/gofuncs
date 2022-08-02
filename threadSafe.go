package gofuncs

import (
	"github.com/jinzhu/copier"
	"sync"
)

type ThreadSafeMap[K comparable, V any] struct {
	mapping map[K]V
	mutex   sync.Mutex
}

func MakeThreadSafeMap[K comparable, V any]() *ThreadSafeMap[K, V] {
	tsMap := &ThreadSafeMap[K, V]{}
	tsMap.mapping = map[K]V{}
	return tsMap
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

func (threadMap *ThreadSafeMap[K, V]) RangeOverCopy(elementHandler func(key K, value V)) {
	threadMap.mutex.Lock()

	copiedMap := map[K]V{}
	err := copier.Copy(&copiedMap, &(threadMap.mapping))
	if err != nil {
		Panic("", err)
	}

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
