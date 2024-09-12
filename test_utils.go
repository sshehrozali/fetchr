package main

import (
	"os"
)

func setMockCliInput(mockInput string) {
	// Open pipe for read-write I/O
	r, w, err := os.Pipe()
	if err != nil {
		return
	}

	w.Write([]byte(mockInput))	// write mock data in stdin
	w.Close()

	// Replace stdin with pipe read-end
	os.Stdin = r
}

