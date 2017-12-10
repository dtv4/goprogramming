package main

import (
	"fmt"
)

//interface

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func test_byteCounter() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	fmt.Fprintf(&c, "%s %s", "byte", "counter")
	fmt.Println(c)
}
