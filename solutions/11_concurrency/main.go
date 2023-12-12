package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	f1, err := os.Open("./a.txt")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	f2, err := os.Open("./b.txt")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	resultFile, err := os.OpenFile("./c.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer resultFile.Close()

	var wg sync.WaitGroup
	errorsCh := make(chan error)
	in := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		readFileByLine(f1, in, errorsCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		readFileByLine(f2, in, errorsCh)
	}()

	go func() {
		wg.Wait()
		close(in)
		close(errorsCh)
	}()

	for line := range in {
		if _, err := resultFile.WriteString(line + "\n"); err != nil {
			panic(err)
		}
	}

	for err := range errorsCh {
		fmt.Println(err)
	}
}

func readFileByLine(file *os.File, out chan string, errorsCh chan error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out <- scanner.Text()
	}

	if scanner.Err() != nil {
		errorsCh <- scanner.Err()
	}
}
