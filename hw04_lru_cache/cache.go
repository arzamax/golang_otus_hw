package hw04_lru_cache //nolint:golint,stylecheck

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

type dict struct {
	data map[Key]*cacheItem
}

func (d *dict) read(key Key) (*cacheItem, bool) {
	res, ok := d.data[key]
	return res, ok
}

func (d *dict) write(key Key, value interface{}, queue *LinkedList, capacity int) (*cacheItem, bool) {
	var (
		res *cacheItem
		ok  bool
	)
	res, ok = d.read(key)

	if ok {
		i := res
		i.value = value
		queue.MoveToFront(i.key)
	} else {
		if capacity == queue.Length {
			delete(d.data, queue.Tail.Value.(Key))
			queue.Remove(queue.Tail)
		}

		d.data[key] = &cacheItem{
			value: value,
			key:   queue.PushFront(key),
		}
		res = d.data[key]
	}

	return res, ok
}

type lruCache struct {
	capacity int
	queue    *LinkedList
	dict     dict
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	_, ok := cache.dict.write(key, value, cache.queue, cache.capacity)
	return ok
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := cache.dict.read(key); ok {
		cache.queue.MoveToFront(i.key)
		return i.value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.dict = dict{data: make(map[Key]*cacheItem)}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		dict:     dict{data: make(map[Key]*cacheItem)},
	}
}
