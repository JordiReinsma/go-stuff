package cache

import (
	"container/list"
	"fmt"
	"math/rand"
)

// Entry is a key/value pair of the cache content
type Entry(type K comparable, V interface{}) struct {
	K
	V
}

// Cache manages least recently used items, with a variable
// size limit and evict policies. It is designed to load the
// value of a key if the key is not present in the cache
type Cache(type K comparable, V interface{}) struct {
	items   map[K]*list.Element
	lruList *list.List
	loader  func(K) (V, error)
	maxSize int
}

// Returns a pointer to a new Cache, given a size limit and
// a loader function. Panics if the size is <= 1
func New(type K comparable, V interface{})(size int, loader func(K) (V, error)) *Cache(K, V) {
	if size <= 1 {
		panic("Provided size is too small")
	}
	return &Cache(K, V){
		items:   make(map[K]*list.Element, size + 1),
		lruList: list.New(),
		loader:  loader,
		maxSize: size,
	}
}

func (c *Cache(K, V)) add(key K) (*list.Element, error) {
	val, err := c.loader(key)
	if err != nil {
		return nil, err
	}
	element := c.lruList.PushFront(&Entry(K, V){key, val})
	c.items[key] = element
	if c.lruList.Len() > c.maxSize {
		c.remove()
	}
	return element, nil
}

func (c *Cache(K, V)) remove() {
	element := c.lruList.Back()
	c.lruList.Remove(element)
	kv := element.Value.(*Entry(K, V))
	delete(c.items, kv.K)
}

// Get returns the value given a key. If the key is not
// present in the cache, it loads its value and removes
// the least recently used item if the cache has reached
// its size limit, to give place to the new key
func (c *Cache(K, V)) Get(key K) (V, error) {
	element, ok := c.items[key]
	if !ok {
		var err error
		element, err = c.add(key)
		if err != nil {
			var zero V
			return zero, err
		}
	} else {
		c.lruList.MoveToFront(element)
	}
	return element.Value.(*Entry(K, V)).V, nil
}

// Clear erases all items from the cache
func (c *Cache(K, V)) Clear() {
	for key := range c.items {
		delete(c.items, key)
	}
	c.lruList.Init()
}

// Entries returns the cache entries, sorted by least
// recently used
func (c *Cache(K, V)) Entries() []Entry(K, V) {
	entries := make([](Entry(K, V)), len(c.items))
	i := 0
	for kv := c.lruList.Back(); kv != nil; kv = kv.Prev() {
		entries[i] = *kv.Value.(*Entry(K, V))
		i++
	}
	return entries
}
