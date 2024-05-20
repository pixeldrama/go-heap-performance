# Heap vs Stack Performance

## Call By Value and Call by Reference

- Call by value default for datatypes: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr,
  float32, float64, string, bool, byte, rune, Array, Structs
- Slices and Maps are Called by Reference by default
- Receiver functions

## Stack vs Heap

- Call by value does not use the heap, no GC involved here
- Call by reference could use the heap, it is not always the case

## Execution

```shell
go build -gcflags '-m -l' # -m print decisions from compile, -l avoid inlining
go test -bench=. -benchmem 
```

## Some Take Away
- Call by seems to be faster in our artificial example
- API-Design mutability should not be mixed
- No absense value is possible (like `nil`), typical solution is to wrap it into a struct: 

```go
type Counter struct {
    valid boolean
    i int
}
```

## References

- [Benchmarking Generics in Go](https://programmingpercy.tech/blog/benchmarking-generics-in-go/)
- [Golang Pass by value vs Pass by reference](https://david-yappeter.medium.com/golang-pass-by-value-vs-pass-by-reference-e48aac8b2716)
- [Understanding Allocations in Go](https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d)
- [When to use Pointers](https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac)