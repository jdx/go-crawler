package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteURL_AlreadyHasHost(t *testing.T) {
	url, err := absoluteURL("https://abc/foo/bar", "https://monzo.com/abc")
	assert.NoError(t, err)
	assert.Equal(t, "https://abc/foo/bar", url)
}

func TestAbsoluteURL_NoHost(t *testing.T) {
	url, err := absoluteURL("/foo/bar", "https://monzo.com/abc")
	assert.NoError(t, err)
	assert.Equal(t, "https://monzo.com/foo/bar", url)
}

func TestAbsoluteURL_Relative(t *testing.T) {
	url, err := absoluteURL("foo/bar", "https://monzo.com/abc")
	assert.NoError(t, err)
	assert.Equal(t, "https://monzo.com/abc/foo/bar", url)
}
