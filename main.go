package main

import (
	"fmt"
	"os"
	"filedownloader/network"
)

// GLOBALs
var exit = os.Exit







func promptInput() string {
	var url string

	fmt.Print("Enter download URL: ")
	fmt.Scanln(&url)

	if url == "" {
		fmt.Println("\nProvided download URL is empty.")
		exit(1)
	}

	return url
}

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
