package hw04lrucache

import "sync"

type Key string

var mx sync.Mutex

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(k Key, value interface{}) bool {
	mx.Lock()
	defer mx.Unlock()
	v, ok := c.items[k]
	if ok {
		v.Value = value
		c.queue.MoveToFront(v)
		c.items[k] = c.queue.Front()
	} else {
		if c.capacity == c.queue.Len() {
			for key, value := range c.items {
				if value == c.queue.Back() {
					delete(c.items, key)
				}
			}
			c.queue.Remove(c.queue.Back())
		}
		c.queue.PushFront(value)
		c.items[k] = c.queue.Front()
	}
	return ok
}

func (c *lruCache) Get(k Key) (interface{}, bool) {
	mx.Lock()
	defer mx.Unlock()
	if v, ok := c.items[k]; ok {
		c.queue.MoveToFront(v)
		return v.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	for k := range c.items {
		delete(c.items, k)
	}
}

// type cacheItem struct {
// 	key   string
// 	value interface{}
// }

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
