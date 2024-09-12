package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptInput(t *testing.T) {
	mockUrl := "www.example.com/download\n"

    originalStdin := os.Stdin

    // Open pipe for read-write
    r, w, err := os.Pipe()
    if err != nil {
        return
    }

    w.Write([]byte(mockUrl))
    w.Close()

    // Replace stdin with read-end
    os.Stdin = r

    actualUrl := promptInput()

    assert.Equal(t, "www.example.com/download", actualUrl, "Expected equals actual")

    // Restore original OS Stdin
    os.Stdin = originalStdin
}

func TestDownloadFile(t *testing.T) {
	// xyz logic
}

func TestSaveLocally(t *testing.T) {
	// xyz logic
}