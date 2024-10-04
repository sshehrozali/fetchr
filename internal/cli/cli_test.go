package cli

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
)

func Test_PromptInput_IfUrlIsNotEmptyThenReturnUrl(t *testing.T) {
	os.Args = []string{"cmd", "-url", "dummy_download_url"}

	_, err := PromptInput()

	assert.Error(t, err, "Error occured")
}