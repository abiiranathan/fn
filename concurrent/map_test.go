package concurrent_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/abiiranathan/fn/concurrent"
)

func TestMapNew(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	if m == nil {
		t.Error("want a valid map, got nil")
	}
}

func TestMapSet(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)

	if m.Len() != 2 {
		t.Errorf("want 2 items, got %d", m.Len())
	}
}

func TestMapGet(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)

	v, ok := m.Get("one")
	if !ok {
		t.Error("want key to be found, got not found")
	}
	if v != 1 {
		t.Errorf("want 1, got %d", v)
	}

	_, ok = m.Get("three")
	if ok {
		t.Error("want key to be not found, got found")
	}
}

func TestMapDelete(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)

	m.Delete("one")
	if m.Len() != 1 {
		t.Errorf("want 1 item, got %d", m.Len())
	}

	_, ok := m.Get("one")
	if ok {
		t.Error("want key to be not found, got found")
	}
}

func TestMapRange(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)

	var keys []string
	m.Range(func(k string, v int) bool {
		keys = append(keys, k)
		return true
	})

	if len(keys) != 2 {
		t.Errorf("want 2 keys, got %d", len(keys))
	}
}

func TestMapKeys(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)

	keys := m.Keys()
	if len(keys) != 2 {
		t.Errorf("want 2 keys, got %d", len(keys))
	}

	for _, k := range keys {
		_, ok := m.Get(k)
		if !ok {
			t.Errorf("want key %s to be found, got not found", k)
		}
	}
}

func TestMapLen(t *testing.T) {
	m := concurrent.NewMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)

	if m.Len() != 2 {
		t.Errorf("want 2 items, got %d", m.Len())
	}
}

func TestConcurrentAccess(t *testing.T) {
	m := concurrent.NewMap[string, int]()

	// insert 10 values concurrently
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			m.Set(fmt.Sprintf("%d", i), i)
		}(i)
	}
	wg.Wait()

	if m.Len() != 10 {
		t.Errorf("want 10 items, got %d", m.Len())
	}
}
