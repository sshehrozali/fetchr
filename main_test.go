package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptInputIfUrlIsValidThenReturn(t *testing.T) {
	// Given
	originalStdin := os.Stdin // make a copy of original stdin before execution
	mockUrl := "www.example.com/download\n"
	setMockCliInput(mockUrl)

	// When
	actualUrl := promptInput()

	// Then
	assert.Equal(t, "www.example.com/download", actualUrl, "Expected equals actual")

	os.Stdin = originalStdin // restore original stdin
}

func TestPromptInputIfUrlIsEmptyThenExitwithCode1(t *testing.T) {
	// arrange
    originalStdin := os.Stdin
    originalStdout := os.Stdout
	exit = mockExit // replace original os.Exit() with mockExit()
	exitCode = 0    // default exit code
	mockInvalidUrl := "\n"
	setMockCliInput(mockInvalidUrl)
    stdReadOutput, stdWriteOutput := captureStdOutput()

    // act
	promptInput()

    // assert
    assert.Equal(t, 1, exitCode, "Expected exit code 1 for empty URL")
    assertStdOutput(stdReadOutput, stdWriteOutput, "Provided download URL is em", t)

    // cleanup
    os.Stdin = originalStdin
    os.Stdout = originalStdout
}

func TestDownloadFile(t *testing.T) {
	// xyz logic
}

func TestSaveLocally(t *testing.T) {
	// xyz logic
}
