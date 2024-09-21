package cli

import (
	"filedownloader/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptInputIfUrlIsNotEmptyThenReturnUrl(t *testing.T) {
	tests.SetMockCliInput("download_url")

	url, _ := PromptInput()

	assert.Equal(t, "download_url", url)
}