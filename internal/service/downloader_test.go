package service

import (
	"errors"
	"filedownloader/internal/models"
	"filedownloader/shared/tests"
	"fmt"

	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fakeDownloadUrl = "url"
var fakeDownloadedData = []byte("fake binary data")

func Test_Download_IfHttpStatusIs200ThenReturnDownloadResult(t *testing.T) {

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

func Test_Download_IfHttpStatusIsNot200ThenReturnError(t *testing.T) {
	mockHttpClient := new(tests.MockHttpClient)
	mockFileStorage := new(tests.MockFileStorage)

	mockHttpResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(bytes.NewBufferString("internal server error")),
		Header:     http.Header{"Content-Type": []string{"text/plain"}},
	}

	mockHttpClient.On("HttpGet", fakeDownloadUrl).Return(mockHttpResponse, nil)

	subject := NewDownloader(mockHttpClient, mockFileStorage)

	_, err := subject.Download(fakeDownloadUrl)

	assert.Error(t, err, "Error occured")

	mockHttpClient.AssertExpectations(t)
}

func Test_Save_IfFileIsSavedLocallyThenPrintFileSizeOnCli(t *testing.T) {
	fakeFileSizeInBytes := 10008232
	mockHttpClient := new(tests.MockHttpClient)
	mockFileStorage := new(tests.MockFileStorage)

	mockFileStorage.On("SaveLocally", "downloaded_file.txt", fakeDownloadedData).Return(fakeFileSizeInBytes, nil)

	rOut, wOut := tests.CaptureStdOutput()

	subject := NewDownloader(mockHttpClient, mockFileStorage)

	err := subject.Save(models.DownloadResult{
		Data:     fakeDownloadedData,
		MimeType: "text/plain",
	})

	assert.NoError(t, err, "No error occured")
	tests.AssertStdOutput(rOut, wOut, fmt.Sprintf("Downloaded file size: %d bytes", fakeFileSizeInBytes), t)

	mockFileStorage.AssertExpectations(t)
}

func Test_Save_IfFileFailedToSaveLocallyThenReturnError(t *testing.T) {
	mockHttpClient := new(tests.MockHttpClient)
	mockFileStorage := new(tests.MockFileStorage)

	mockFileStorage.On("SaveLocally", "downloaded_file.txt", fakeDownloadedData).Return(0, errors.New("mock error message"))

	subject := NewDownloader(mockHttpClient, mockFileStorage)

	err := subject.Save(models.DownloadResult{
		Data:     fakeDownloadedData,
		MimeType: "text/plain",
	})

	assert.Error(t, err, "Error occured")

	mockFileStorage.AssertExpectations(t)
}
