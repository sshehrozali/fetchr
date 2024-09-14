package main

import (
	"os"
	"testing"
	"net/http"
	"bytes"
	"io"
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
	assertStdOutput(stdReadOutput, stdWriteOutput, "Provided download URL is empty.", t)

	// cleanup
	os.Stdin = originalStdin
	os.Stdout = originalStdout
}

func TestDownloadFile(t *testing.T) {
	// arrange
	mockHttpClient := new(MockHttpClient)
	mockResponse := &http.Response{
        StatusCode: http.StatusOK,
        Body:       io.NopCloser(bytes.NewBufferString("mock body")),
        Header:     http.Header{"Content-Type": []string{"text/plain"}},
    }

	mockHttpClient.On("Get", "http://example.com").Return(mockResponse, nil)	// stub the http call

	// act
	result := downloadFile(mockHttpClient, "http://example.com")

	// assert
	assert.Equal(t, string("mock body"), result.Data)
	assert.Equal(t, "text/plain", result.MimeType)

	mockHttpClient.AssertExpectations(t) // verify mock was called
}

func TestSaveLocally(t *testing.T) {
	// xyz logic
}
