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
	keyMap   map[*ListItem]Key
	items    map[Key]*ListItem
}

func (cache lruCache) Set(key Key, value interface{}) bool {
	if _, ok := cache.items[key]; !ok {
		if cache.queue.Len() == cache.capacity {
			lastElement := cache.queue.Back()
			cache.queue.Remove(lastElement)
			delete(cache.items, cache.keyMap[lastElement])
			delete(cache.keyMap, lastElement)
		}
		cache.items[key] = cache.queue.PushFront(value)
		cache.keyMap[cache.items[key]] = key
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
		keyMap:   make(map[*ListItem]Key, capacity),
		items:    make(map[Key]*ListItem, capacity),
	}
}
