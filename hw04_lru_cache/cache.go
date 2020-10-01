package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	value interface{}
	key   *LinkedListItem
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    *LinkedList
	dict     map[Key]*cacheItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.Lock()
	defer cache.Unlock()

	res, ok := cache.dict[key]

	if ok {
		i := res
		i.value = value
		cache.queue.MoveToFront(i.key)
	} else {
		if cache.capacity == cache.queue.Length {
			delete(cache.dict, cache.queue.Tail.Value.(Key))
			cache.queue.Remove(cache.queue.Tail)
		}

		cache.dict[key] = &cacheItem{
			value: value,
			key:   cache.queue.PushFront(key),
		}
	}

	return ok
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.Lock()
	defer cache.Unlock()

	if i, ok := cache.dict[key]; ok {
		cache.queue.MoveToFront(i.key)
		return i.value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.dict = make(map[Key]*cacheItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		dict:     make(map[Key]*cacheItem),
	}
}
