package main

import (
	"github.com/stretchr/testify/mock"
	"net/http"
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