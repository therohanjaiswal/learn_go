package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	ReadSomething()

	if err := ReadSomethingSimple(); err != nil {
		fmt.Printf("error occured: %s", err)
	}
}

func ReadSomethingSimple() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	// defer func obeys stack property i.e., LIFO
	// defer are used to close all the open connections
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is okay")
		} else if p != nil {
			panic(errors.New("an unexpected error occurred but we do not want to recover"))
		}
	}()

	defer func() {
		println("before for-loop")
	}()

	for {
		val, readerErr := r.Read([]byte("text does nothing"))
		if readerErr == io.EOF {
			fmt.Println("finished reading file, breaking out of the loop.\n")
			break
		} else if readerErr != nil {
			err = readerErr
			return
		}
		println(val)
	}

	defer func() {
		println("after for-loop")
	}()

	return nil
}

type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("panic: something catastrophic occurred.")

func (sr *SimpleReader) Read(p []byte) (n int, err error) {
	if sr.count == 2 {
		panic(errors.New("Another error."))
	}
	if sr.count > 3 {
		return 0, io.EOF
	}
	sr.count++
	return sr.count, nil
}

func (sr *SimpleReader) Close() error {
	println("closing reader")
	return nil
}

func ReadSomething() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	value, err := r.Read([]byte("test Something"))
	if err != nil {
		fmt.Printf("An error occurred: %s.\n", err)
		return err
	}

	println(value)
	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}
