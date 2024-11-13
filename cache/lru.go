package cache

import (
	"container/list"
	"fmt"
)

// Кэш на основе LRU
type LRUCache struct {
	capacity int                      // Максимальная емкость кэша
	cache    map[string]*list.Element // Хранение данных в виде хэша
	order    *list.List               // Двусвязный список для порядка использования
}

// Элемент кэша
type entry struct {
	key   string
	value string
}

// Создает новый LRU-кэш с заданной емкостью
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element, capacity),
		order:    list.New(),
	}
}

// Получает значение из кэша
func (c *LRUCache) Get(key string) (string, bool) {
	if elem, found := c.cache[key]; found {
		c.order.MoveToFront(elem) // Перемещение в начало как самый недавно использованный
		return elem.Value.(*entry).value, true
	}
	return "", false // Элемент не найден
}

// Добавляет значение в кэш
func (c *LRUCache) Put(key, value string) {
	if elem, found := c.cache[key]; found {
		// Обновление существующего элемента
		c.order.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	// Добавление нового элемента
	if c.order.Len() >= c.capacity {
		// Если кэш заполнен, удаляем последний элемент (наименее использованный)
		oldest := c.order.Back()
		if oldest != nil {
			c.order.Remove(oldest)
			delete(c.cache, oldest.Value.(*entry).key)
		}
	}

	// Добавление нового элемента в начало
	newEntry := &entry{key, value}
	elem := c.order.PushFront(newEntry)
	c.cache[key] = elem
}

// Пример использования
func main() {
	cache := NewLRUCache(3)
	cache.Put("A", "value1")
	cache.Put("B", "value2")
	cache.Put("C", "value3")

	fmt.Println(cache.Get("A")) // Output: value1, true

	cache.Put("D", "value4") // Вытесняет "B" так как он менее использован

	_, found := cache.Get("B") // Output: false, так как "B" был вытеснен
	fmt.Println("B found:", found)

	fmt.Println(cache.Get("C")) // Output: value3, true
}
