package fn_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/abiiranathan/fn"
)

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(v int) bool { return v%2 == 0 }
	want := []int{2, 4}
	got := fn.Filter(s, f)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(v int) int { return v * v }
	want := []int{1, 4, 9, 16, 25}
	got := fn.Map(s, f)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestForEach(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	doubled := make([]int, 0, len(s))
	f := func(v int) {
		doubled = append(doubled, v*2)
	}
	fn.ForEach(s, f)

	want := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(want, doubled) {
		t.Errorf("want %v, got %v", want, doubled)
	}
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(acc, v int) int { return acc + v }
	want := 15
	got := fn.Reduce(s, f, 0)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	// test reduce with initial value and structs
	type person struct {
		Name string
		Age  int
	}

	s2 := []person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	fn2 := func(acc int, p person) int { return acc + p.Age }
	want2 := 90
	got2 := fn.Reduce(s2, fn2, 0)
	if want2 != got2 {
		t.Errorf("want %v, got %v", want2, got2)
	}

	// advanced use case of reduce
	fn3 := func(acc map[string]int, p person) map[string]int {
		acc[p.Name] = p.Age
		return acc
	}

	want3 := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	got3 := fn.Reduce(s2, fn3, make(map[string]int))
	if !reflect.DeepEqual(want3, got3) {
		t.Errorf("want %v, got %v", want3, got3)
	}
}

func TestConcat(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	want := []int{1, 2, 3, 4, 5}
	got := fn.Concat(s1, s2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestIndexOf(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	elem := 3
	want := 2
	got := fn.IndexOf(s, elem)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	// test element not found
	elem = 6
	want = -1
	got = fn.IndexOf(s, elem)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDistinct(t *testing.T) {
	s := []int{1, 2, 2, 3, 3, 3}
	want := []int{1, 2, 3}
	got := fn.Distinct(s)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDistinctFunc(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	s := []person{
		{"Alice", 25},
		{"Bob", 30},
		{"Alice", 25},
		{"John", 40},
		{"John", 40},
	}

	f := func(p person) interface{} { return p.Name }
	want := []person{
		{"Alice", 25},
		{"Bob", 30},
		{"John", 40},
	}

	got := fn.DistinctFunc(s, f)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestChunk(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chunkSize := 3
	want := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	got := fn.Chunk(s, chunkSize)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

	// check uneven chunk size
	chunkSize = 4
	want = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9}}
	got = fn.Chunk(s, chunkSize)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestChunkEmpty(t *testing.T) {
	s := []int{}
	chunkSize := 3
	want := [][]int{}
	got := fn.Chunk(s, chunkSize)

	// check empty slice
	if len(got) != 0 {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestChunkSmall(t *testing.T) {
	s := []int{1, 2}
	chunkSize := 3
	want := [][]int{{1, 2}}
	got := fn.Chunk(s, chunkSize)

	// check small slice
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestPartition(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	is_even := func(v int) bool { return v%2 == 0 }
	wantYes := []int{2, 4}
	wantNo := []int{1, 3, 5}
	evens, odd := fn.Partition(s, is_even)
	if !reflect.DeepEqual(wantYes, evens) || !reflect.DeepEqual(wantNo, odd) {
		t.Errorf("want %v, %v, got %v, %v", wantYes, wantNo, evens, odd)
	}
}

func TestFlatten(t *testing.T) {
	s := [][]int{{1, 2, 3}, {4, 5}}
	want := []int{1, 2, 3, 4, 5}
	got := fn.Flatten(s)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	want := []int{5, 4, 3, 2, 1}
	fn.Reverse(s)
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}
}

func TestShuffle(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	fn.Shuffle(s)
	if len(s) != 5 {
		t.Errorf("want %v, got %v", 5, len(s))
	}
}

func TestShuffleEmpty(t *testing.T) {
	s := []int{}
	fn.Shuffle(s)
	if len(s) != 0 {
		t.Errorf("want %v, got %v", 0, len(s))
	}
}

func TestTake(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	n := 3
	want := []int{1, 2, 3}
	got := fn.Take(s, n)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

	// if n is greater than the length of s
	n = 6
	want = []int{1, 2, 3, 4, 5}
	got = fn.Take(s, n)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestTakeWhile(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(v int) bool { return v < 3 }
	want := []int{1, 2}
	got := fn.TakeWhile(s, f)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

	// if all elements satisfy the predicate
	f = func(v int) bool { return v < 6 }
	want = []int{1, 2, 3, 4, 5}
	got = fn.TakeWhile(s, f)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestDrop(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	n := 3
	want := []int{4, 5}
	got := fn.Drop(s, n)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

	// if n is greater than the length of s
	n = 6
	want = nil
	got = fn.Drop(s, n)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestCount(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(v int) bool { return v%2 == 0 }
	want := 2
	got := fn.Count(s, f)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAll(t *testing.T) {
	s := []int{2, 4, 6, 8, 10}
	f := func(v int) bool { return v%2 == 0 }
	want := true
	got := fn.All(s, f)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	// test when not all elements satisfy the predicate
	s = []int{2, 4, 6, 7, 10}
	want = false
	got = fn.All(s, f)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestAny(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(v int) bool { return v%2 == 0 }
	want := true
	got := fn.Any(s, f)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}

	// test when no elements satisfy the predicate
	s = []int{1, 3, 5, 7, 9}
	want = false
	got = fn.Any(s, f)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestZip(t *testing.T) {
	s1 := []int{10, 20, 30}
	s2 := []string{"Abiira", "Dan", "Joseph"}

	type person struct {
		Age  int
		Name string
	}

	f := func(age int, name string) person {
		return person{age, name}
	}

	want := []person{
		{10, "Abiira"},
		{20, "Dan"},
		{30, "Joseph"},
	}

	got := fn.Zip(s1, s2, f)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}

	// zip with unequal length
	s1 = []int{10, 20}
	s2 = []string{"Abiira", "Dan", "Joseph"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	fn.Zip(s1, s2, f)

}

func TestZipWithIndex(t *testing.T) {
	s := []string{"Abiira", "Dan", "Joseph"}

	f := func(i int, name string) string {
		return fmt.Sprintf("%s is at index %d", name, i)
	}

	want := []string{
		"Abiira is at index 0",
		"Dan is at index 1",
		"Joseph is at index 2",
	}

	got := fn.ZipWithIndex(s, f)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestZipShortest(t *testing.T) {
	s1 := []int{10, 20, 30}
	s2 := []string{"Abiira", "Dan"}

	type person struct {
		Age  int
		Name string
	}

	f := func(age int, name string) person {
		return person{age, name}
	}

	want := []person{
		{10, "Abiira"},
		{20, "Dan"},
	}

	got := fn.ZipShortest(s1, s2, f)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestRotateLeft(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	fn.RotateLeft(s, 2)
	want := []int{3, 4, 5, 1, 2}
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}

	// test when positions is greater than the length of s
	fn.RotateLeft(s, 10)
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}

	// pos <2
	s = []int{1}
	fn.RotateLeft(s, 2)
	want = []int{1}
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}
}

func TestRotateRight(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	fn.RotateRight(s, 2)
	want := []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}

	// test when positions is greater than the length of s
	fn.RotateRight(s, 10)
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}

	// pos <2
	s = []int{1}
	fn.RotateRight(s, 2)
	want = []int{1}
	if !reflect.DeepEqual(want, s) {
		t.Errorf("want %v, got %v", want, s)
	}
}
