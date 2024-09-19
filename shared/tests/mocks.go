package shared

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

var exitCode int

func mockExit(code int) {
	exitCode = code
}

// Mock Http client
type MockHttpClient struct {
	mock.Mock
}
func (mock *MockHttpClient) Get(url string) (*http.Response, error) {
	args := mock.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}