package cli

import (
	"errors"
	"flag"
	"strings"
)

func PromptInput() (string, error) {

	url := flag.String("url", "", "URL of the file to download")
	flag.Parse()

	if *url == "" {
		return "", errors.New("please provide a URL using the -url flag")
	}

	if !strings.HasPrefix(*url, "https://") && !strings.HasPrefix(*url, "http://") {
		return "", errors.New("the URL must contain http/s")
	}

	return *url, nil
}