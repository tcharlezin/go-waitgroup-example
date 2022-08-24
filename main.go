package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	messages := []string{
		"Hello universe!",
		"Hello cosmos!",
		"Hello world!",
	}

	for _, value := range messages {
		wg.Add(1)
		go updateMessage(value)
		wg.Wait()

		printMessage()
	}

}
