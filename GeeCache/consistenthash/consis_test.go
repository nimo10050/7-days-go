package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"testing"
)

// HashFunc 定义一个 hash 算法
type HashFunc func(data []byte) uint32

// HashRing 定义哈希环
type HashRing struct {
	replicas int            // 虚拟节点的数量
	keys     []int          // 哈希环
	hashFunc HashFunc       // 定义的哈希算法
	hashMap  map[int]string // 虚拟节点与真实节点的映射关系 key 为虚拟节点的哈希值， 值是真实节点的名称
}

// New creates a Map instance
func NewHashRing(replicas int, fn HashFunc) *HashRing {
	m := &HashRing{
		replicas: replicas,
		hashFunc: fn,
		hashMap:  make(map[int]string),
	}
	if m.hashFunc == nil {
		m.hashFunc = crc32.ChecksumIEEE
	}
	return m
}

// replicas 为虚拟节点的数量, key 为真实节点拿来计算 hash 值的， 比如机器IP、名称等
func (m *HashRing) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func TestHashRing(t *testing.T) {
	hash := NewHashRing(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// Given the above hash function, this will give replicas with "hashes":
	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	hash.Add("6", "4", "2")

	fmt.Println(hash.keys)
}
