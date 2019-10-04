package main

import (
	"testing"
)

func Test_CounterAdd(t *testing.T) {
	c := New()
	c.Add("foo")
	c.Add("foo")
	c.Add("bar")

	n := c.Get("foo")
	if n != 2 {
		t.Logf(`expected 2 "foo", got %d`, n)
		t.Fail()
	}

	n = c.Get("bar")
	if n != 1 {
		t.Logf(`expected 1 "bar", got %d`, n)
		t.Fail()
	}
}
