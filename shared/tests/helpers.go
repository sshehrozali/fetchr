package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetMockCliInput(mockInput string) {
	// Open pipe for read-write I/O
	r, w, _ := os.Pipe()

	w.Write([]byte(mockInput)) // write mock data in stdin
	w.Close()

	// Replace stdin with pipe read-end
	os.Stdin = r
}

func CaptureStdOutput() (*os.File, *os.File) {
	// Capture os.Stdout output to verify printed messages
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut // write stdout in pipe

	return rOut, wOut
}

func AssertStdOutput(rOut *os.File, wOut *os.File, expected string, t *testing.T) {
	// Close and reset os.Stdout so we can read the output
	wOut.Close()
	var buf bytes.Buffer
	buf.ReadFrom(rOut)

	assert.Contains(t, buf.String(), expected)
}

func StartMockWebServer(endpoint string, responseData string) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseData))
	})

	return httptest.NewServer(handler)
}
