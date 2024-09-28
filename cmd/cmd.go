package cmd

import (
	"fetchr/internal/cli"
	"fetchr/internal/network"
	"fetchr/internal/service"
	"fetchr/internal/storage"
	"log"
)

func Run() error {
	url, cliErr := cli.PromptInput()

	if cliErr != nil {
		log.Fatal(cliErr.Error())
	}

	httpClient := &network.HttpClient{}
	fileStorage := &storage.FileStorage{}

	downloader := service.NewDownloader(httpClient, fileStorage)

	result, downloadErr := downloader.Download(url)

	if downloadErr != nil {
		log.Fatal(downloadErr.Error())
	}

	saveErr := downloader.Save(result)

	if saveErr != nil {
		return saveErr
	}

	return nil
}
