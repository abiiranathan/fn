# fn

fn is a Go package that provides a set of functional programming tools for Go. It is inspired by the functional programming features of languages like Haskell, Scala, Rust and JavaScript.

## Installation

```bash
go get github.com/abiiranathan/fn
```

## API
Functions:

- `Filter`: Returns a new slice containing only the elements that satisfy the predicate.
- `Map`: Returns a new slice containing the results of applying a function to each element.
- `Reduce`: Accumulates the elements of a slice by applying a function.
- `Concat`: Concatenates two slices.
- `IndexOf`: Returns the index of the first occurrence of an element.
- `Distinct`: Returns a new slice containing only the unique elements of a slice.
- `DistinctFunc`: Returns a new slice containing only the unique elements of a slice based on a key function.
- `Chunk`: Returns a new slice containing slices of a specified size.
- `Shuffle`: Randomizes the order of elements in a slice.
- `Partition`: Partitions a slice into two based on a predicate function.
- `Flatten`: Flattens a slice of slices into a single slice.
- `Reverse`: Reverses the elements of a slice in place.
- `Take`: Returns the first n elements of a slice.
- `TakeWhile`: Returns elements from the beginning of a slice as long as a condition is true.
- `Drop`: Returns a new slice without the first n elements.
- `Count`: Returns the number of elements in a slice that satisfy a condition.
- `All`: Returns true if all elements in a slice satisfy a condition.
- `Any`: Returns true if any element in a slice satisfies a condition.
- `Zip`: Applies a function to pairs of elements from two slices.
- `ZipShortest`: Applies a function to pairs of elements from two slices, stopping at the shorter slice.
- `ZipWithIndex`: Applies a function to elements of a slice and their index.
- `RotateLeft`: Rotates the elements of a slice to the left.
- `RotateRight`: Rotates the elements of a slice to the right.



## Usage

1. Filter

```go
s := []int{1, 2, 3, 4, 5}
is_even := func(v int) bool { return v%2 == 0 }
got := fn.Filter(s, is_even)
fmt.Println(got) // [2 4]
```

2. Map

```go
s := []int{1, 2, 3, 4, 5}
double := func(v int) int { return v * 2 }
got := fn.Map(s, double)
fmt.Println(got) // [2 4 6 8 10]
```

3. Reduce

```go
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
got2 := fn.Reduce(s2, 0, f2)
fmt.Println(got2) // 90


// More advanced example
fn3 := func(acc map[string]int, p person) map[string]int {
    acc[p.Name] = p.Age
    return acc
}
got3 := fn.Reduce(s2, make(map[string]int), f3)

fmt.Println(got3) // map[Alice:25 Bob:30 Charlie:35]
```

4. Concat

```go
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}
got := fn.Concat(s1, s2)
fmt.Println(got) // [1 2 3 4 5 6]
```

5. IndexOf

```go
s := []int{1, 2, 3, 4, 5}
got := fn.IndexOf(s, 3)
fmt.Println(got) // 2
```

6. Distinct

```go
s := []int{1, 2, 2, 3, 3, 3}
got := fn.Distinct(s)
fmt.Println(got) // [1 2 3]
```

7. DistinctFunc

```go
type person struct {
    Name string
    Age  int
}

s := []person{
    {"Bob", 30},
    {"Alice", 25},
    {"Bob", 30},
    {"Alice", 35},
    {"Jane", 12},
}

key := func(p person) string { return p.Name }
got := fn.DistinctFunc(s, key)
fmt.Println(got) // [{Bob 30} {Alice 25} {Jane 12}]
```

8. Chunk

```go
s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
got := fn.Chunk(s, 3)
fmt.Println(got) // [[1 2 3] [4 5 6] [7 8 9]]
```

9. Shuffle

```go
s := []int{1, 2, 3, 4, 5}
got := fn.Shuffle(s)
// Output is non-deterministic
fmt.Println(got)
``` 

10. Partition

```go
s := []int{1, 2, 3, 4, 5}
is_even := func(v int) bool { return v%2 == 0 }
even, odd := fn.Partition(s, is_even)
fmt.Println(got) // [2 4], [1 3 5]
```

11. Flatten

```go
s := [][]int{{1, 2}, {3, 4}, {5, 6}}
got := fn.Flatten(s)
fmt.Println(got) // [1 2 3 4 5 6]
```

12. Reverse

```go
s := []int{1, 2, 3, 4, 5}
fn.Reverse(s)
fmt.Println(s) // [5 4 3 2 1]
```

13. Take

```go
s := []int{1, 2, 3, 4, 5}
got := fn.Take(s, 3)
fmt.Println(got) // [1 2 3]
```

14. TakeWhile

```go
s := []int{1, 2, 3, 4, 5}
is_less_than_4 := func(v int) bool { return v < 4 }
got := fn.TakeWhile(s, is_less_than_4)
fmt.Println(got) // [1 2 3]
```

15. Drop

```go
s := []int{1, 2, 3, 4, 5}
got := fn.Drop(s, 3)
fmt.Println(got) // [4 5]
```

16. Count

```go
s := []int{1, 2, 3, 4, 5}
is_even := func(v int) bool { return v%2 == 0 }
got := fn.Count(s, is_even)
fmt.Println(got) // 2
```

17. All

```go
s := []int{1, 2, 3, 4, 5}
is_less_than_10 := func(v int) bool { return v < 10 }
got := fn.All(s, is_less_than_10)
fmt.Println(got) // true
```

18. Any

```go
s := []int{1, 2, 3, 4, 5}
is_even := func(v int) bool { return v%2 == 0 }
got := fn.Any(s, is_even)
fmt.Println(got) // true
```

19. Zip

```go
s1 := []int{1, 2, 3}
s2 := []string{"a", "b", "c"}
f := func(a int, b string) string { return fmt.Sprintf("%d%s", a, b) }
got := fn.Zip(s1, s2, f)
fmt.Println(got) // [1a 2b 3c]
```

20. ZipShortest

```go
s1 := []int{1, 2, 3}
s2 := []string{"a", "b"}
f := func(a int, b string) string { return fmt.Sprintf("%d%s", a, b) }
got := fn.ZipShortest(s1, s2, f)
fmt.Println(got) // [1a 2b]
```

21. ZipWithIndex

```go
s := []string{"a", "b", "c"}

f := func(v string, i int) string { return fmt.Sprintf("%d%s", i, v) }
got := fn.ZipWithIndex(s, f)
fmt.Println(got) // [0a 1b 2c]
```

22. RotateLeft

```go
s := []int{1, 2, 3, 4, 5}
got := fn.RotateLeft(s, 2)
fmt.Println(got) // [3 4 5 1 2]
```

23. RotateRight

```go
s := []int{1, 2, 3, 4, 5}
got := fn.RotateRight(s, 2)
fmt.Println(got) // [4 5 1 2 3]
```

## License

MIT


