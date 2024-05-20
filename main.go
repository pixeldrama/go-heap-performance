package main

import "fmt"

func main() {
	count := 0

	inc(count)
	fmt.Printf("Add by value: %d\n", count)

	incPointer(&count)
	fmt.Printf("Add by reference %d\n", count)
}

func inc(count int) {
	count++
}

func incPointer(count *int) {
	*count++
}

/*func main() {
	counter := Counter{i: 0}

	counter.counterInc()
	fmt.Printf("Add by value %d\n", counter.i)

	counter.counterIncPointer()
	fmt.Printf("Add by value %d\n", counter.i)
}

type Counter struct {
	i int
}

func (c Counter) counterInc() {
	c.i++
}

func (c *Counter) counterIncPointer() {
	c.i++
}*/

type Person struct {
	Name string
	Age  int
}

//go:noinline
func NewPersonCopy(name string, age int) Person {
	return Person{Name: name, Age: age}
}

//go:noinline
func NewPersonPointer(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
