package tests

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

var ExitCode int

func MockExit(code int) {
	ExitCode = code
}

// Mock Http client
type MockHttpClient struct {
	mock.Mock
}
func (mock *MockHttpClient) HttpGet(url string) (*http.Response, error) {
	args := mock.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

// Mock File storage
type MockFileStorage struct {
	mock.Mock
}
func (mock *MockFileStorage) SaveLocally(fileName string, data []byte) (int, error) {
	args := mock.Called(fileName, data)
	return args.Get(0).(int), args.Error(1)
}