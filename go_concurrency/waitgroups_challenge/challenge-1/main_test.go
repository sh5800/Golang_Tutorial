package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("epsilon", &wg)

	wg.Wait()

	if msg != "epsilon" {
		t.Errorf("Expected 'epsilon' but got %s", msg)
	}
}

func TestPrintMessge(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	msg = "alpha"

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut	

	if output != "alpha\n" {
		t.Errorf("Expected 'alpha' but got %s", output)
	}
}

func TestMain(t *testing.T){
	stdOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut	

	if !strings.Contains(output,"Hello, universe!"){
		t.Errorf("Expected 'Hello, universe!' but got %s", output)
	}

	if !strings.Contains(output,"Hello, cosmos!"){
		t.Errorf("Expected 'Hello, cosmos!' but got %s", output)
	}

	if !strings.Contains(output,"Hello, world!"){
		t.Errorf("Expected 'Hello, world!' but got %s", output)
	}
}
