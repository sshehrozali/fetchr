package cli

import (
	"errors"
	"flag"
	"strings"
)

func PromptInput() (string, error) {

	url := flag.String("url", "", "URL of the file to download")
	// output := flag.String("o", "output_file", "Output file path (default: output_file)")
	flag.Parse()

	if *url == "" {
		return "", errors.New("please provide a URL using the -url flag")
	}

	if !strings.HasPrefix(*url, "https://") {
		return "", errors.New("The URL must contain https://")
	}

	return *url, nil
}