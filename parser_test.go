package linkparser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLinks(t *testing.T) {
	html := `<a href="/test">Test Link</a>`
	reader := strings.NewReader(html)

	links, err := ParseReader(reader)
	assert.NoError(t, err)
	assert.Len(t, links, 1)
	assert.Equal(t, "/test", links[0].Href)
	assert.Equal(t, "Test Link", links[0].Content)
}
