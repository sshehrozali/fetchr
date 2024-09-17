package main

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"
)

var exitCode int

func mockExit(code int) {
	exitCode = code
}

// Mock network client
type MockHttpClient struct {
	mock.Mock
}
func (mock *MockHttpClient) Get(url string) (*http.Response, error) {
	args := mock.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

// Mock file writer
type MockFileWriter struct {
	WriteError bool
}
func (w *MockFileWriter) Write(b []byte) (int, error) {
	if w.WriteError {
		return 0, errors.New("error writing bytes in file")
	}

	return len(b), nil
}