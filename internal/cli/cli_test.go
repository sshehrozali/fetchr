package cli

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
)

var url string

func Test_PromptInput_IfUrlIsNotEmptyThenReturnUrl(t *testing.T) {
	os.Args = []string{"cmd", "-url", "dummy_download_url"}

	url, _ := PromptInput()

	assert.Equal(t, "dummy_download_url", url)
}

// func Test_PromptInput_IfUrlIsEmptyThenReturnError(t *testing.T) {
// 	flag.StringVar(&url, "url", "", "URL to download")
// 	os.Args = []string{"cmd", "-url", ""}

// 	_, err := PromptInput()

// 	assert.Error(t, err, "Error occured")
// }