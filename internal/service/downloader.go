package service

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"filedownloader/internal/models"
	"filedownloader/internal/utils"
)

type NetworkClient interface {
	httpGet(url string) (*http.Response, error)
	// ftp
	// other protocols...
}

type Storage interface {
	Save(data []byte) (int, error)
	// Load
	// fetch, etc..
}

type Downloader struct {
	networkClient NetworkClient
	storage       Storage
}

func NewDownloader(networkClient NetworkClient, storage Storage) *Downloader {
	return &Downloader{
		networkClient: networkClient,
		storage:       storage,
	}
}

func (d *Downloader) download(url string) (models.DownloadResult, error) {
	fmt.Println("\nFetching data from the server...")

	// then start loader
	done := make(chan bool)
	utils.

	// For HTTP/s URL, perform blocking GET
	// response, err := d.networkClient.httpGet(url)

	// Stop the loader
	// done <- true
	// close(done)

	if err != nil {
		return models.DownloadResult{}, errors.New("error while performing HTTP GET request")
	}

	if response.StatusCode != http.StatusOK {
		return models.DownloadResult{}, errors.New("download request failed with HTTP status: " + string(response.StatusCode))
	}

	mimeType := response.Header.Get(("Content-Type"))
	fmt.Println("\nDetected content-type: ", mimeType)

	body, err := io.ReadAll(response.Body) // Parse body into byte[]

	if err != nil {
		return models.DownloadResult{}, errors.New("error while decoding response body")
	}

	return models.DownloadResult{
		Data:     body,
		MimeType: mimeType,
	}, nil
}
