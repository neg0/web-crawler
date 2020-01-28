package assets

import (
	"cuvva/pkg/adapter/html"
)

type extraction struct {
	Document html.Finder
}

func newAssetsExtractor(HTMLContent []byte) (*extraction, error) {
	htmlReader, err := html.NewReader(HTMLContent)
	if err != nil {
		return nil, err
	}

	return &extraction{
		Document: htmlReader,
	}, nil
}

func (e extraction) Title() string {
	titles := e.Document.Find("title")
	if len(titles) > 0 {
		return e.Document.Find("title")[0].Text()
	}
	return ""
}

func (e extraction) CSS() []string {
	var cssCollection []string
	cssFiles := e.Document.Find("link")
	for _, css := range cssFiles {
		if css.Attr("rel") == "stylesheet" {
			cssCollection = append(cssCollection, css.Attr("href"))
		}
	}
	return cssCollection
}

func (e extraction) JS() []string {
	var scriptsCollection []string
	scripts := e.Document.Find("script")
	for _, script := range scripts {
		if len(script.Attr("src")) == 0 {
			scriptsCollection = append(scriptsCollection, script.Text())
			continue
		}
		scriptsCollection = append(scriptsCollection, script.Attr("src"))
	}
	return scriptsCollection
}

func (e extraction) IMG() []string {
	var imgCollection []string
	images := e.Document.Find("img")
	for _, img := range images {
		if len(img.Attr("src")) != 0 {
			imgCollection = append(imgCollection, img.Attr("src"))
		}
	}
	return imgCollection
}
