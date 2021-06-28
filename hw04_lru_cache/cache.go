package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mx       sync.Mutex
}

type cacheItem struct {
	key   string
	value interface{}
}

func (c *lruCache) Set(k Key, v interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()

	item := &cacheItem{key: string(k), value: v}

	if el, ok := c.items[k]; ok {
		c.queue.MoveToFront(el)
		el.Value.(*cacheItem).value = v
		return true
	}

	if c.queue.Len() == c.capacity {
		r := c.queue.Back().Value.(*cacheItem).key
		delete(c.items, Key(r))
		c.queue.Remove(c.queue.Back())
	}

	e := c.queue.PushFront(item)
	c.items[k] = e

	return false
}

func (c *lruCache) Get(k Key) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()

	if el, ok := c.items[k]; ok {
		c.queue.MoveToFront(el)
		return el.Value.(*cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
