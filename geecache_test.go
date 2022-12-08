package geecache

import (
	"fmt"
	"log"
	"testing"
)

func TestGetGroup(t *testing.T) {
	group1 := NewGroup("topic1", 100, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			return nil, fmt.Errorf("%s not exist", key)
		}),
	)

	fmt.Println(&group1.mainCache.lru)
	fmt.Println(&group1)

	group2 := NewGroup("topic2", 100, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			return nil, fmt.Errorf("%s not exist", key)
		}),
	)
	fmt.Println(&group2)
	fmt.Println(&group2.mainCache.lru)
}
