package lru

import "container/list"

type Cache struct {
	maxBytes  int64
	nBytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	onEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// New 创建 Cache 对象
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

// Add 添加或者更新操作
func (c *Cache) Add(key string, value Value) {
	// 如果缓存对象已经存在
	if ele, ok := c.cache[key]; ok {
		// 移到队列的最前面
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		// 重新计算缓存大小
		c.nBytes = c.nBytes - int64(kv.value.Len()) + int64(value.Len())
		kv.value = value
	} else {
		// 把元素放到队列的最前面
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nBytes = c.nBytes + int64(len(key)) + int64(value.Len())
	}

	// 触发缓存淘汰
	if c.maxBytes >= 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

// Get 查找操作
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest 移除操作
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes = c.nBytes - int64(len(kv.key)) - int64(kv.value.Len())
		if c.onEvicted != nil {
			c.onEvicted(kv.key, kv.value)
		}
	}
}
