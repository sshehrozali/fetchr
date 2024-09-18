package main

import (
	"fmt"
	"os"
	"filedownloader/network"
)











func run() {
	url := promptInput()

	httpClient := &network.DefaultHttpClient{}
	downloadResult := downloadFile(httpClient, url)

	fileWriter := &file.DefaultFileWriter{}
	saveLocally(downloadResult)
}

func main() {
	run()
}
