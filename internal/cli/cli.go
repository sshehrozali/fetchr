package cli

import (
	"errors"
	"flag"
)

func PromptInput() (string, error) {

	url := flag.String("url", "", "URL of the file to download")
	// output := flag.String("o", "output_file", "Output file path (default: output_file)")
	flag.Parse()

<<<<<<< HEAD
	// Error Handling

	if url == "" {
		return "", errors.New("url can't be empty")
	}

	if len(url) < 8 || url[:8] != "https://" {
		return "", errors.New("The URL must contain https://")
	}

	return url, nil

=======
	if *url == "" {
		return "", errors.New("please provide a URL using the -url flag")
	}

	return *url, nil
>>>>>>> 7a3fa8c967bd63bd6281ddef57f1b6a33612a015
}

// Sample URL -> https://filesampleshub.com/format/image/jpg
