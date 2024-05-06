package concurrent_test

import (
	"testing"

	"github.com/abiiranathan/fn/concurrent"
)

func TestSetNew(t *testing.T) {
	m := concurrent.NewSet[string]()
	if m == nil {
		t.Error("want a valid map, got nil")
	}

	if m.Len() != 0 {
		t.Errorf("want 0 items, got %d", m.Len())
	}
}

func TestSetAdd(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	if m.Len() != 2 {
		t.Errorf("want 2 items, got %d", m.Len())
	}

}

func TestSetContains(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	if !m.Contains("one") {
		t.Error("want key to be found, got not found")
	}

	if m.Contains("three") {
		t.Error("want key to be not found, got found")
	}
}

func TestSetRemove(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	m.Remove("one")
	if m.Len() != 1 {
		t.Errorf("want 1 item, got %d", m.Len())
	}

	if m.Contains("one") {
		t.Error("want key to be not found, got found")
	}
}

func TestSetValues(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	values := m.Values()
	if len(values) != 2 {
		t.Errorf("want 2 items, got %d", len(values))
	}

	// Order is not guaranteed in a map, so we can't check for the exact order.
	// if values[0] != "one" || values[1] != "two" {
	// 	t.Errorf("want [one two], got %v", values)
	// }
}

func TestSetUnion(t *testing.T) {
	m1 := concurrent.NewSet[string]()
	m1.Add("one")
	m1.Add("two")

	m2 := concurrent.NewSet[string]()
	m2.Add("two")
	m2.Add("three")

	m3 := m1.Union(m2)
	if m3.Len() != 3 {
		t.Errorf("want 3 items, got %d", m3.Len())
	}

	if !m3.Contains("one") || !m3.Contains("two") || !m3.Contains("three") {
		t.Errorf("want [one two three], got %v", m3.Values())
	}
}

func TestSetIntersection(t *testing.T) {
	m1 := concurrent.NewSet[string]()
	m1.Add("one")
	m1.Add("two")

	m2 := concurrent.NewSet[string]()
	m2.Add("two")
	m2.Add("three")

	m3 := m1.Intersection(m2)
	if m3.Len() != 1 {
		t.Errorf("want 1 items, got %d", m3.Len())
	}

	if !m3.Contains("two") {
		t.Errorf("want [two], got %v", m3.Values())
	}
}

func TestSetDifference(t *testing.T) {
	m1 := concurrent.NewSet[string]()
	m1.Add("one")
	m1.Add("two")

	m2 := concurrent.NewSet[string]()
	m2.Add("two")
	m2.Add("three")

	m3 := m1.Difference(m2)
	if m3.Len() != 1 {
		t.Errorf("want 1 items, got %d", m3.Len())
	}

	if !m3.Contains("one") {
		t.Errorf("want [one], got %v", m3.Values())
	}
}

func TestSetSymmetricDifference(t *testing.T) {
	m1 := concurrent.NewSet[string]()
	m1.Add("one")
	m1.Add("two")

	m2 := concurrent.NewSet[string]()
	m2.Add("two")
	m2.Add("three")

	m3 := m1.SymmetricDifference(m2)
	if m3.Len() != 2 {
		t.Errorf("want 2 items, got %d", m3.Len())
	}

	if !m3.Contains("one") || !m3.Contains("three") {
		t.Errorf("want [one three], got %v", m3.Values())
	}
}

func TestSetIsSubset(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	subset := concurrent.NewSet[string]()
	subset.Add("one")

	if !subset.IsSubset(m) {
		t.Error("want subset, got not a subset")
	}
}

func TestSetIsSuperset(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	superset := concurrent.NewSet[string]()
	superset.Add("one")
	superset.Add("two")
	superset.Add("three")

	if !superset.IsSuperset(m) {
		t.Error("want superset, got not a superset")
	}
}

func TestSetClone(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	clone := m.Clone()
	if clone.Len() != 2 {
		t.Errorf("want 2 items, got %d", clone.Len())
	}

	if !clone.Contains("one") || !clone.Contains("two") {
		t.Errorf("want [one two], got %v", clone.Values())
	}
}

func TestSetClear(t *testing.T) {
	m := concurrent.NewSet[string]()
	m.Add("one")
	m.Add("two")

	m.Clear()
	if m.Len() != 0 {
		t.Errorf("want 0 items, got %d", m.Len())
	}
}
