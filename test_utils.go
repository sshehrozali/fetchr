package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setMockCliInput(mockInput string) {
	// Open pipe for read-write I/O
	r, w, _ := os.Pipe()

	w.Write([]byte(mockInput)) // write mock data in stdin
	w.Close()

	// Replace stdin with pipe read-end
	os.Stdin = r
}

func captureStdOutput() (*os.File, *os.File) {
	// Capture os.Stdout output to verify printed messages
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut // redirect all output to the pipe

	return rOut, wOut
}

func assertStdOutput(rOut *os.File, wOut *os.File, expected string, t *testing.T) {
	// Close and reset os.Stdout so we can read the output
	wOut.Close()
	var buf bytes.Buffer
	buf.ReadFrom(rOut)

	assert.Contains(t, buf.String(), expected)
}
