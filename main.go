package main

import "fmt"

// New initializes a new Counter and its backing map.
func New() Counter {
	var c Counter
	c.members = make(map[string]int)

	return c
}

// Counter uses a map[<type>]int to count how many instances of a value have been observed.
// Inspired by Python's collections.Counter type.
type Counter struct {
	members map[string]int
}

// Add accepts a value and increments its observation count by 1.
func (c *Counter) Add(m string) {
	c.members[m]++
}

// Get returns the number of times a value has been observed.
func (c *Counter) Get(m string) int {
	n, ok := c.members[m]
	if !ok {
		return 0
	}

	return n
}

// Most returns the n most frequently observed values, as a Counter.
// This is to faciliate further filtering of the provided data.
func (c *Counter) Most(n int) Counter {
	var r Counter

	return r
}

// Least returns the n least frequently observed values, as a Counter.
// This is to faciliate further filtering of the provided data.
func (c *Counter) Least(n int) Counter {
	var r Counter

	return r
}

func main() {
	fmt.Println("vim-go")
}
