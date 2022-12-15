package geecache

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplitN(t *testing.T) {
	// path: /test/1
	p := "/_geecache_/"
	s := "/_geecache_/groupname/key"
	n := strings.SplitN(s[len(p):], "/", 2)
	fmt.Println(n)
}
