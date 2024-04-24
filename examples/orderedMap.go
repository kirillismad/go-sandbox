package examples

import (
	"container/list"
	"fmt"
)

func DemoOrderedMap() {
	demoOrderedMap()
}

func demoOrderedMap() {
	orderedMap := NewOrderedMap[string, int]()
	orderedMap.Set("ONE", 1)
	orderedMap.Set("FOUR", 4)
	orderedMap.Set("TWO", 2)
	orderedMap.Set("FIVE", 5)
	orderedMap.Set("THREE", 3)

	orderedMap.Set("ONE", -1)
	orderedMap.Delete("FOUR")

	for iter := orderedMap.getIterator(); iter.HasNext(); {
		key, value := iter.Next()
		fmt.Println(key, value)
	}
}

type OrderedMap[K comparable, V any] struct {
	keys         *list.List
	keysElements map[K]*list.Element
	values       map[K]V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:         list.New(),
		keysElements: make(map[K]*list.Element),
		values:       make(map[K]V),
	}
}

// Set method doesn't change the order of map keys
func (m *OrderedMap[K, V]) Set(key K, value V) {
	_, exists := m.values[key]
	if !exists {
		m.keysElements[key] = m.keys.PushBack(key)
	}
	m.values[key] = value
}

// Set method moves key to the end
// func (m *OrderedMap[K, V]) Set(key K, value V) {
// 	element, exists := m.keyelements[key]
// 	if !exists {
// 		m.keyelements[key] = m.keys.PushBack(key)
// 	} else {
// 		m.keys.MoveToBack(element)
// 	}
// 	m.values[key] = value
// }

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	value, exists := m.values[key]
	return value, exists
}

func (m *OrderedMap[K, V]) Delete(key K) {
	element, exists := m.keysElements[key]
	if exists {
		delete(m.values, key)
		delete(m.keysElements, key)
		m.keys.Remove(element)
	}

}
func (m *OrderedMap[K, V]) getIterator() *iterator[K, V] {
	return &iterator[K, V]{orderedMap: m, current: m.keys.Front()}
}

type iterator[K comparable, V any] struct {
	orderedMap *OrderedMap[K, V]
	current    *list.Element
}

func (i *iterator[K, V]) HasNext() bool {
	return i.current != nil
}

func (i *iterator[K, V]) Next() (K, V) {
	key := i.current.Value.(K)
	value := i.orderedMap.values[key]

	i.current = i.current.Next()

	return key, value
}
