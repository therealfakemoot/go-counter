package counter

import (
	"fmt"
	"sort"
)

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

// Reset resets all observations, replacing the backing map with a new one.
func (c *Counter) Reset() {
	c.members = make(map[string]int)
}

func (c *Counter) sorted() []string {
	var s []string
	for k := range c.members {
		s = append(s, k)
	}

	sort.Slice(s, func(i, j int) bool {
		return c.Get(s[i]) > c.Get(s[j])
	})

	return s
}

// Get returns the number of times a value has been observed.
func (c *Counter) Get(m string) int {
	n, ok := c.members[m]
	if !ok {
		return 0
	}

	return n
}

// Set sets the observation count for a given value to the provided int.
func (c *Counter) Set(m string, n int) {
	c.members[m] = n
}

// Most returns the n most frequently observed values, as a Counter.
// This is to faciliate further filtering of the provided data.
func (c *Counter) Most(n int) Counter {
	sorted := c.sorted()
	r := New()

	for i := 0; i < len(sorted)-1; i++ {
		r.Set(sorted[i], c.Get(sorted[i]))
	}

	return r
}

// Least returns the n least frequently observed values, as a Counter.
// This is to faciliate further filtering of the provided data.
func (c *Counter) Least(n int) Counter {
	sorted := c.sorted()
	r := New()

	for i := len(sorted) - 1; i >= 0; i-- {
		r.Set(sorted[i], c.Get(sorted[i]))
	}

	return r
}

func main() {
	fmt.Println("vim-go")
}