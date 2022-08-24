package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	wg.Add(1)
	updateMessage("teste")
	wg.Wait()

	if msg != "teste" {
		t.Error("Expected teste but didn't found")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Gopher"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Gopher") {
		t.Error("Expected to find Gopher")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello universe!") {
		t.Error("Expected to find Hello universe!")
	}

	if !strings.Contains(output, "Hello cosmos!") {
		t.Error("Expected to find Hello cosmos!")
	}

	if !strings.Contains(output, "Hello world!") {
		t.Error("Expected to find Hello world!")
	}
}
