package utils

import (
	"fmt"
	"time"
)

func FileExtensionRetriever(mimeType string) string {
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

func Loader(done chan bool) {
	symbols := []string{"|", "/", "-", "\\"}

	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%s", symbols[i%len(symbols)])
			i++
			time.Sleep(100 * time.Millisecond)

		}
	}
}
