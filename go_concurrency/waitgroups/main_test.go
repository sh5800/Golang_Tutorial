package main

import (
	"io"
	"os"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r,w,_ := os.Pipe()

	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)

	printSomething("Hello World", &wg)

	wg.Wait()

	_ = w.Close()

	result,_ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if output != "Hello World\n" {
		t.Errorf("Expected 'Hello World' but got %s", output)
	}

}
