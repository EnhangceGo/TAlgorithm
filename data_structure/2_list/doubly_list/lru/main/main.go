package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int                   // 缓存容量
	cache    map[int]*list.Element // 存储每个键对应的双向链表节点
	list     *list.List            // 双向链表，存储缓存的键值对
}

type CacheItem struct {
	key   int // 键
	value int // 值
}

// 创建并返回一个新的 LRUCache 对象
func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// 获取缓存中指定键的值，如果不存在返回 -1，否则将键对应的节点移动到链表头部，并返回节点的值
func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		// 将节点移动到链表头部，表示最近使用过
		this.list.MoveToFront(elem)
		// 返回节点的值
		return elem.Value.(*CacheItem).value
	}
	// 键不存在，返回 -1
	return -1
}

// 添加或更新缓存中的键值对，如果键已经存在，更新值并将节点移动到链表头部。否则，如果缓存已满，删除链表尾部的节点并从 cache 中删除对应的键值对。然后创建新的 CacheItem 对象，并将其插入到链表头部和 cache 中。
func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		// 键已经存在，更新值并将节点移动到链表头部，表示最近使用过
		elem.Value.(*CacheItem).value = value
		this.list.MoveToFront(elem)
	} else {
		// 键不存在，创建新的 CacheItem 对象
		item := &CacheItem{key: key, value: value}
		elem := this.list.PushFront(item)
		// 将键对应的节点插入到 cache 中，表示已经被缓存
		this.cache[key] = elem
		// 如果缓存已满，删除链表尾部的节点并从 cache 中删除对应的键值对
		if this.list.Len() > this.capacity {
			lastElem := this.list.Back()
			delete(this.cache, lastElem.Value.(*CacheItem).key)
			this.list.Remove(lastElem)
		}
	}
}

func main() {
	cache := NewLRUCache(2)

	// Add key-value pairs to the cache
	cache.Put(1, 1)
	cache.Put(2, 2)

	// Retrieve a value from the cache
	fmt.Println(cache.Get(1)) // Output: A

	// Add another key-value pair to the cache
	cache.Put(3, 3)

	// The oldest item (key=2) should be evicted from the cache
	fmt.Println(cache.Get(2)) // Output: <nil>
}
