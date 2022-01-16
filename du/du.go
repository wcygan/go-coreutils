package du

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

// Run executes the disk usage program
// source: https://www.gopl.io/ chapter 8
func Run(dir string, out io.Writer) error {
	go shutDownWhenInputIsDetected()

	fileSizes := make(chan int64)
	var wg sync.WaitGroup

	wg.Add(1)
	go walkDir(dir, &wg, fileSizes)

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	tick := time.Tick(500 * time.Millisecond)
	var numberOfFiles, numberOfBytes int64

loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				// drain fileSizes to prevent goroutine leak
			}
			return nil
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			numberOfFiles++
			numberOfBytes += size
		case <-tick:
			printDiskUsage(numberOfFiles, numberOfBytes, out)
		}
	}

	printDiskUsage(numberOfFiles, numberOfBytes, out)
	return nil
}

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(numberOfFiles, numberOfBytes int64, out io.Writer) {
	output := fmt.Sprintf("%d files  %.1f GB", numberOfFiles, float64(numberOfBytes)/1e9)
	_, err := fmt.Fprintln(out, output)
	if err != nil {
		return
	}
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()
	if canceled() {
		return
	}

	for _, entry := range directoryEntries(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDirectory := filepath.Join(dir, entry.Name())
			go walkDir(subDirectory, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var semaphore = make(chan struct{}, 30)

func directoryEntries(dir string) []os.FileInfo {
	select {
	case semaphore <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-semaphore }()

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}

// shutDownWhenInputIsDetected shuts the program down when input is detected from os.Stdin
func shutDownWhenInputIsDetected() {
	// Block while waiting for input
	os.Stdin.Read(make([]byte, 1))
	close(done)
}
