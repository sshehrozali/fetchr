package cli

import (
	"errors"
	"fmt"
)

func PromptInput() (string, error) {
	var url string

	fmt.Print("Enter download URL: ")
	fmt.Scanln(&url)

	// Error Handling

	if url == "" {
		return "", errors.New("url can't be empty")
	}

	if len(url) < 8 || url[:8] != "https://" { 
		return "", errors.New("The URL must contain https://")
	}

	return url, nil

}


// Sample URL --> https://filesampleshub.com/format/image/jpg