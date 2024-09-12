package main

import (

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptInputIfUrlIsValidThenReturn(t *testing.T) {
    // Given
	originalStdin := os.Stdin   // make a copy of original stdin before execution
    mockUrl := "www.example.com/download\n"
    setMockCliInput(mockUrl)

    // When
	actualUrl := promptInput()

    // Then
	assert.Equal(t, "www.example.com/download", actualUrl, "Expected equals actual")

	os.Stdin = originalStdin // restore original stdin
}

func TestPromptInputIfUrlIsEmptyThenExitwithCode1(t *testing.T) {
	exit = mockExit
	exitCode = 0 // reset exit code before execution

	mockUrl := "\n"


	// Open pipe for read-write
	r, w, err := os.Pipe()
	if err != nil {
		return
	}

	w.Write([]byte(mockUrl))
	w.Close()

	// Replace stdin with read-end
	os.Stdin = r

	// // Capture output to check the printed error message
	// var stdOutputBuf bytes.Buffer
	// fmt.

	// actualUrl := promptInput()

}

func TestDownloadFile(t *testing.T) {
	// xyz logic
}

func TestSaveLocally(t *testing.T) {
	// xyz logic
}
