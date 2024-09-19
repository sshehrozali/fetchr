package service

import (
	"bytes"
	"filedownloader/shared/tests"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

var fakeDownloadUrl = "url"
var fakeDownloadedData = []byte("fake binary data")

func TestDownloadIfHttpStatusIs200ThenReturnDownloadResult(t *testing.T) {

	mockHttpClient := new(tests.MockHttpClient)
	mockFileStorage := new(tests.MockFileStorage)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(string(fakeDownloadedData))),
		Header:     http.Header{"Content-Type": []string{"text/plain"}},
	}

	mockHttpClient.On("HttpGet", fakeDownloadUrl).Return(mockHttpResponse, nil)

	subject := NewDownloader(mockHttpClient, mockFileStorage)

	result, _ := subject.Download(fakeDownloadUrl)

	assert.Equal(t, fakeDownloadedData, result.Data)
	assert.Equal(t, "text/plain", result.MimeType)

	mockHttpClient.AssertExpectations(t)
}
