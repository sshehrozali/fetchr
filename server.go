package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fileExtensionRetriever(mimeType string) string {
	commonMIMEtypes := map[string]string{
		// Text types
		"text/plain":      ".txt",
		"text/html":       ".html",
		"text/css":        ".css",
		"text/javascript": ".js",

		// Image types
		"image/jpeg":    ".jpeg",
		"image/png":     ".png",
		"image/gif":     ".gif",
		"image/svg+xml": ".svg",

		// Audio types
		"audio/mpeg": ".mp3",
		"audio/wav":  ".wav",
		"audio/ogg":  ".ogg",

		// Video types
		"video/mp4":  ".mp4",
		"video/webm": ".webm",
		"video/ogg":  ".ogv",

		// Application types
		"application/json":                  ".json",
		"application/xml":                   ".xml",
		"application/pdf":                   ".pdf",
		"application/zip":                   ".zip",
		"application/octet-stream":          ".bin",
		"application/x-www-form-urlencoded": ".url",

		// Multipart types
		"multipart/form-data": ".multpart",
	}

	fileExtension := commonMIMEtypes[mimeType]
	return fileExtension
}

func main() {

	// Perform HTTP GET request
	// Ensure status is 200OK and no error
	// Check file MIME type
	// Decode body in bytes -> make sure no errors
	// Transform, write the file locally

	var url string

	fmt.Print("Enter download URL: ")
	fmt.Scanln(&url)

	if url == "" {
		fmt.Println("\nProvided download URL is empty.")
		os.Exit(1)
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("\nError while performing GET request.")
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Request failed with HTTP status: ", response.StatusCode)
		os.Exit(1)
	}

	mimeType := response.Header.Get(("Content-Type"))
	fmt.Println("\nDetected content-type: ", mimeType)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while reading response body.")
		os.Exit(1)
	}

	fileName := "downloaded_file" + fileExtensionRetriever(mimeType)
	file, err := os.Create(fileName)
	defer file.Close()

	bytesWritten, err := file.Write(body)

	if err != nil {
		fmt.Println("Error copying bytes locally.")
		os.Exit(1)
	}

	if bytesWritten <= 0 {
		fmt.Println("0 bytes copied.")
		os.Exit(1)
	}

	fmt.Printf("Downloaded file size: %d bytes\n", bytesWritten)
	fmt.Println(fmt.Sprintf("File downloaded successfully. Check your current local directory %s", fileName))
	os.Exit(0)
}
