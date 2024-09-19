package service

import ("testing"
"filedownloader/shared/tests")

func TestDownloadIfHttpStatusIs200ThenReturnDownloadResult(t *testing.T) {

	mockHttpClient := new(tests.MockHttpClient)
	mockFileStorage := new(tests.MockFileStorage)

	subject := NewDownloader(mockHttpClient, mockFileStorage)
}