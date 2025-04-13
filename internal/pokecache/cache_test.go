package pokecache

import (
	"testing"
	"time"
)

func TestGetAdd(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://test.example",
			val: []byte("data"),
		},
		{
			key: "http://test.example/path",
			val: []byte("more data"),
		},
	}

	for _, c := range cases {
		cache := NewCache(interval)
		cache.Add(c.key, c.val)
		val, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key")
			return
		}
		if string(val) != string(c.val) {
			t.Errorf("expected val to be %s", c.val)
			return
		}
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Second
	const wait = interval + 5*time.Millisecond
	key := "https://test.example"
	val := []byte("data")

	cache := NewCache(interval)
	cache.Add(key, val)

	time.Sleep(wait)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
