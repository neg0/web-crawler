package crawler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSiteMapCrawler(t *testing.T) {
	sut := NewSiteMapCrawler()

	actual, err := sut.Handle("https://cuvva.com")

	assert.NoError(t, err)
	assert.Len(t, actual, 157)
}
