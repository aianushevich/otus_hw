package hw04lrucache

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
	item := cache.items[key]
	if item == nil {
		cache.queue.PushFront(value)
		cache.items[key] = cache.queue.Front()
		if cache.queue.Len() > cache.capacity {
			lastElement := cache.queue.Back()
			cache.queue.Remove(lastElement)
			for k, v := range cache.items {
				if v.Value == lastElement.Value {
					delete(cache.items, k)
				}
			}
		}
		return false
	} else {
		item.Value = value
		cache.queue.MoveToFront(item)
		cache.items[key] = cache.queue.Front()
		return true
	}
}

func (cache lruCache) Get(key Key) (interface{}, bool) {
	if len(cache.items) == 0 {
		return nil, false
	}

	item := cache.items[key]
	if item == nil {
		return nil, false
	}
	cache.queue.MoveToFront(item)
	return item.Value, true
}

func (cache lruCache) Clear() {
	cache.items = nil
	cache.queue = nil
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
