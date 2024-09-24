package service

import (
	"errors"
	"fetchr/internal/models"
	"fetchr/internal/utils"
	"fmt"
	"io"
	"net/http"
)

type NetworkClient interface {
	HttpGet(url string) (*http.Response, error)
	// ftp
	// other protocols...
}

type Storage interface {
	SaveLocally(fileName string, data []byte) (int, error)
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

func (d *Downloader) Download(url string) (models.DownloadResult, error) {
	fmt.Println("\nFetching data from the server...")

	// then start loader
	done := make(chan bool)
	go utils.Loader(done)

	// For HTTP/s URL, perform blocking GET
	response, err := d.networkClient.HttpGet(url)

	// Stop the loader
	done <- true
	close(done)

	if err != nil {
		return models.DownloadResult{}, errors.New("error while performing HTTP GET request")
	}

	if response.StatusCode != http.StatusOK {
		return models.DownloadResult{}, fmt.Errorf("download request failed with HTTP status: %d", response.StatusCode)
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

func (d *Downloader) Save(downloadResult models.DownloadResult) error {
	fileName := "downloaded_file" + utils.FileExtensionRetriever(downloadResult.MimeType)

	size, error := d.storage.SaveLocally(fileName, downloadResult.Data)

	if error != nil {
		return errors.New("error saving downloaded file locally")
	}

	fmt.Printf("Downloaded file size: %d bytes\n", size)
	fmt.Printf("File downloaded successfully. Check your current local directory %s\n", fileName)

	return nil
}
