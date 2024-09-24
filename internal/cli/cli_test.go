package cli

import (
	"fetchr/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PromptInput_IfUrlIsNotEmptyThenReturnUrl(t *testing.T) {
	tests.SetMockCliInput("download_url")

	url, _ := PromptInput()

	assert.Equal(t, "download_url", url)
}

func Test_PromptInput_IfUrlIsEmptyThenReturnError(t *testing.T) {
	tests.SetMockCliInput("")

	_, err := PromptInput()

	assert.Error(t, err, "Error occured")
}