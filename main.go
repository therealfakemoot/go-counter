package main

import "fmt"

func New() Counter {
	var c Counter
	c.members = make(map[string]int)

	return c
}

type Counter struct {
	members map[string]int
}

func (c *Counter) Add(m string) {
	c.members[m]++
}

func (c *Counter) Get(m string) int {
	n, ok := c.members[m]
	if !ok {
		return 0
	}

	return n
}

func (c *Counter) Most(n int) Counter {
	var r Counter

	return r
}

func (c *Counter) Least(n int) Counter {
	var r Counter

	return r
}

func main() {
	fmt.Println("vim-go")
}
