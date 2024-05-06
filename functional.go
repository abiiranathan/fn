package fn

import (
	"math/rand"
	"slices"
)

// Filter returns a new slice containing only the elements that satisfy the predicate fn.
func Filter[T any](s []T, fn func(T) bool) []T {
	p := make([]T, 0, len(s))
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return slices.Clip(p)
}

// Map returns a new slice containing the results of applying the function
// fn to each element of the original slice.
func Map[T, U any](s []T, fn func(T) U) []U {
	p := make([]U, len(s))
	for i, v := range s {
		p[i] = fn(v)
	}
	return slices.Clip(p)
}

// Reduce applies the function fn to each element of the slice, accumulating
// the result. The accumulated value is initialized to initial.
func Reduce[T, U any](s []T, fn func(U, T) U, initial U) U {
	p := initial
	for _, v := range s {
		p = fn(p, v)
	}
	return p
}

// Concat returns a new slice containing all the elements of s1
// followed by all the elements of s2.
func Concat[T any](s1, s2 []T) []T {
	return append(s1, s2...)
}

// IndexOf returns the index of the first occurrence of elem in s.
// If elem is not in s, IndexOf returns -1.
func IndexOf[T comparable](s []T, elem T) int {
	for i, v := range s {
		if v == elem {
			return i
		}
	}
	return -1
}

// Distinct returns a new slice containing only the unique elements of s.
func Distinct[T comparable](s []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// DistinctFunc returns a new slice containing only the unique elements of s.
// The function fn is used to determine the uniqueness of each element.
// The function fn should return a value that can be used as a key in a map.
func DistinctFunc[T any, U comparable](s []T, fn func(T) U) []T {
	seen := make(map[interface{}]struct{})
	result := make([]T, 0, len(s))
	for _, v := range s {
		key := fn(v)
		if _, ok := seen[key]; !ok {
			seen[key] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Chunk returns a new slice containing slices of size chunkSize.
// The last slice may have fewer than chunkSize elements.
func Chunk[T any](s []T, chunkSize int) [][]T {
	var chunks [][]T
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}

// Shuffle randomizes the order of elements in s.
func Shuffle[T any](s []T) {
	for i := len(s) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

// Partition returns two slices, the first containing the elements of s
// that satisfy the predicate fn, and the second containing the rest.
func Partition[T any](s []T, fn func(T) bool) (yes, no []T) {
	for _, v := range s {
		if fn(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return
}

// Flatten returns a new slice containing all the elements of the sub-slices in s.
func Flatten[T any](s [][]T) []T {
	var result []T
	for _, subSlice := range s {
		result = append(result, subSlice...)
	}
	return result
}

// Reverse reverses the elements of s in place.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Take returns the first n elements of s. If n is greater than the length of s,
// Take returns s unchanged.
func Take[T any](s []T, n int) []T {
	if n > len(s) {
		return s
	}
	return s[:n]
}

// TakeWhile returns a new slice containing the elements
// that satisfy the predicate fn. The predicate is evaluated
// until the first element that does not satisfy the predicate.
func TakeWhile[T any](s []T, fn func(T) bool) []T {
	for i, v := range s {
		if !fn(v) {
			return s[:i]
		}
	}
	return s
}

// Drop returns a new slice containing all but the first n elements of s.
// This is the opposite of Take. If n is greater than the length of s, Drop returns nil.
func Drop[T any](s []T, n int) []T {
	if n > len(s) {
		return nil
	}
	return s[n:]
}

// Count returns the number of elements in s that satisfy the predicate fn.
func Count[T any](s []T, fn func(T) bool) int {
	count := 0
	for _, v := range s {
		if fn(v) {
			count++
		}
	}
	return count
}

// All returns true if all elements in s satisfy the predicate fn.
func All[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any returns true if at least one element in s satisfies the predicate fn.
func Any[T any](s []T, fn func(T) bool) bool {
	for _, v := range s {
		if fn(v) {
			return true
		}
	}
	return false
}

// Zip returns a new slice containing the result of applying the function fn
// to the elements of s1 and s2. The function fn should take two arguments,
// one from each slice, and return a value.
// s1 and s2 must have the same length.
func Zip[T, U, V any](s1 []T, s2 []U, fn func(T, U) V) []V {
	if len(s1) != len(s2) {
		panic("slices must have the same length")
	}
	result := make([]V, len(s1))
	for i := range s1 {
		result[i] = fn(s1[i], s2[i])
	}
	return result
}

// ZipLongest returns a new slice containing the result of applying the function fn
// to the elements of s1 and s2. The function fn should take two arguments,
// one from each slice, and return a value.
// If the slices have different lengths, the result will have the length of the shorter slice.
func ZipShortest[T, U, V any](s1 []T, s2 []U, fn func(T, U) V) []V {
	var result []V
	minLen := min(len(s1), len(s2))
	for i := 0; i < minLen; i++ {
		result = append(result, fn(s1[i], s2[i]))
	}
	return result
}

// ZipWithIndex returns a new slice containing the result of applying the function fn
// to the elements of s and their index. The function fn should take two arguments,
// the index and the element, and return a value.
func ZipWithIndex[T, U any](s []T, fn func(int, T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = fn(i, v)
	}
	return result
}

// ForEach applies the function fn to each element of s.
func ForEach[T any](s []T, fn func(T)) {
	for _, v := range s {
		fn(v)
	}
}

// RotateLeft rotates the elements of s to the left by positions.
// It does nothing if s has fewer than 2 elements.
// This is useful for working with circular buffers.
func RotateLeft[T any](s []T, positions int) {
	if len(s) < 2 {
		return
	}
	positions = positions % len(s)
	Reverse(s[:positions])
	Reverse(s[positions:])
	Reverse(s)
}

// RotateRight rotates the elements of s to the right by positions.
// It does nothing if s has fewer than 2 elements.
// This is useful for working with circular buffers.
func RotateRight[T any](s []T, positions int) {
	if len(s) < 2 {
		return
	}
	positions = positions % len(s)
	Reverse(s)
	Reverse(s[:positions])
	Reverse(s[positions:])
}
