package lru

import (
	"fmt"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestGet(t *testing.T) {
	lru := New(int64(90), nil)

	key := "zhangsan"

	lru.Add(key, String("123"))

	if value, ok := lru.Get(key); ok {
		fmt.Println("cache value: ", value)
	}

}
