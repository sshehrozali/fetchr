package cmd

import (
	"filedownloader/internal/cli"
	"filedownloader/internal/network"
	"filedownloader/internal/service"
	"filedownloader/internal/storage"
	"os"
)

// GLOBALs
var exit = os.Exit

func Run() error {
	url, cliErr := cli.PromptInput()

	if cliErr != nil {
		println(cliErr.Error())
		exit(1)
	}

	httpClient := &network.HttpClient{}
	fileStorage := &storage.FileStorage{}

	downloader := service.NewDownloader(httpClient, fileStorage)

	result, downloadErr := downloader.Download(url)

	if downloadErr != nil {
		println(downloadErr.Error())
		exit(1)
	}

	saveErr := downloader.Save(result)

	if saveErr != nil {
		return saveErr
	}

	return nil
}
