package sitemap

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSiteMap(t *testing.T) {
	var sut *SiteMap

	t.Run("when is instantiated with a valid XML sitemap", func(t *testing.T) {
		xmlSiteMapMock, err := ioutil.ReadFile("./mocks/sitemap.xml")
		if err != nil {
			t.Fatal("failed to load the mock test case for sitemap", err)
		}

		sut, err = NewSiteMap(xmlSiteMapMock)

		t.Run("should not have any error", func(t *testing.T) {
			t.Log(sut)
			assert.NoError(t, err)
		})

		t.Run("should have all URLs extracted from the XML site map", func(t *testing.T) {
			assert.Len(t, sut.URLs(), 157)
		})
	})

	t.Run("when is instantiated with an invalid XML sitemap", func(t *testing.T) {
		xmlSiteMapMock, err := ioutil.ReadFile("./mocks/invalid_sitemap.xml")
		if err != nil {
			t.Fatal("failed to load the mock test case for sitemap", err)
		}

		sut, err = NewSiteMap(xmlSiteMapMock)

		t.Run("should have an error", func(t *testing.T) {
			assert.Error(t, err)
		})
	})
}
