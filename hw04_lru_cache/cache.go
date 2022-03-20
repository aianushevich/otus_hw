package hw04lrucache

import "fmt"

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
}

func (cache lruCache) Set(key Key, value interface{}) bool {
	if _, ok := cache.items[key]; !ok {
		if cache.queue.Len() == cache.capacity {
			lastElement := cache.queue.Back()
			cache.queue.Remove(lastElement)
			for k, v := range cache.items {
				if v.Value == lastElement.Value {
					delete(cache.items, k)
				}
			}
		}
		cache.items[key] = cache.queue.PushFront(value)
		return false
	}
	cache.queue.Remove(cache.items[key])
	cache.items[key] = cache.queue.PushFront(value)
	return true
}

func (cache lruCache) Get(key Key) (interface{}, bool) {
	item, ok := cache.items[key]
	if !ok {
		return nil, false
	}
	cache.queue.MoveToFront(item)
	cache.items[key] = cache.queue.Front()
	fmt.Println(cache.queue.Front())
	fmt.Println(cache.queue.Back())
	return item.Value, true
}

func (cache *lruCache) Clear() {
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
