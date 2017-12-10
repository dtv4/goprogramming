package main

import (
	"fmt"
	"io"
	"os"
)

type Writer1 int

func (w Writer1) Write(p []byte) (int, error) {
	fmt.Println("Writer1: ", string(p))
	return len(p), nil
}

type Writer2 int

func (w Writer2) Write(p []byte) (int, error) {
	fmt.Println("Writer2: ", string(p))
	return len(p), nil
}

func test_interface() {
	// os.Stdout.Write([]byte("hello"))
	// os.Stdout.Close()

	var w io.Writer
	w = os.Stdout //assign *os.File to interface, OK *os.File has Write method
	w.Write([]byte("hello"))

	var w1 Writer1
	w = w1
	w.Write([]byte("hello"))
}
