package filedownloader

import (
	"fmt"
	"io"
	"os"
	"flag"
	"net/http"
)

func main() {

	// Perform HTTP GET request
	// Ensure status is 200OK and no error
	// Check file MIME type
	// Decode body in bytes -> make sure no errors
	// Transform, write the file locally

	urlFlag := flag.String("url", "", "Download HTTP URl for file")
	flag.Parse()
	url := *urlFlag

	if (url == "") {
		fmt.Println("Provided download URL is empty...")
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error while performing GET request")
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println("Request failed with HTTP status: ", response.StatusCode)
	}

	mimeType := response.Header.Get(("Content-Type"))
	fmt.Println("Detected content-type: ", mimeType)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error while decoding response body")
	}

	commonMIMEtypes := map[string]string{
		// Text types
		"text/plain":        ".txt",
        "text/html":         ".html",
        "text/css":          ".css",
        "text/javascript":   ".js",

        // Image types
        "image/jpeg":        ".jpeg",
        "image/png":         ".png",
        "image/gif":         ".gif",
        "image/svg+xml":     ".svg",

        // Audio types
        "audio/mpeg":        ".mp3",
        "audio/wav":         ".wav",
        "audio/ogg":         ".ogg",

        // Video types
        "video/mp4":         ".mp4",
        "video/webm":        ".webm",
        "video/ogg":         ".ogv",

        // Application types
        "application/json": ".json",
        "application/xml":  ".xml",
        "application/pdf":  ".pdf",
        "application/zip":  ".zip",
        "application/octet-stream": ".bin",
        "application/x-www-form-urlencoded": ".url",
        
        // Multipart types
        "multipart/form-data": ".multpart",
		
	}

	fileExtension := commonMIMEtypes[mimeType]
	fileName := "downloaded_file" + fileExtension
	file, err := os.Create(fileName)
	defer file.Close()
	
	bytesWritten, err := file.Write(body)

	if (err != nil) {
		fmt.Println("Error while writing data in file...")
	}

	if (bytesWritten <= 0) {
		fmt.Println("No bytes written in file...")
	}

	fmt.Println("No of bytes copied: " + string(bytesWritten))
	fmt.Println(fmt.Sprintf("File downloaded successfully. Check your current local directory %s", fileName))

}
