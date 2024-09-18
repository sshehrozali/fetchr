package cli

import (
	"errors"
	"fmt"
)

func PromptInput() (string, error) {
	var url string

	fmt.Print("Enter download URL: ")
	fmt.Scanln(&url)

	if url == "" {
		return "", errors.New("url can't be empty")
	}

	return url, nil
}
