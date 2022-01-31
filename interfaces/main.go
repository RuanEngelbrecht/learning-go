package main

import (
	"bytes"
	"fmt"
)

// Interfaces using a struct

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Interfaces using a type alias

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// Composing interfaces

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// Constructor method
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func main() {
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Go! This is a test"))
	wc.Close()

	r, ok := wc.(*BufferedWriterCloser) // type cast / conversion
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

/*
	Best practices:
	- Use many, small interfaces.
	- Single method interfaces are some of hte most powerful and flexible.
		- io.Reader, io.Writer, interface {}
	- Don't export interfaces for types that will be consumed.
	- Do export interfaces for types that will be used by the package.
	- Design functions and methods to receive interfaces whenever possible.
*/
