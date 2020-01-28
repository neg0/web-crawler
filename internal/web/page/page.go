package page

import (
	"cuvva/internal/web/page/assets"
	"encoding/json"
	"fmt"
	"net/url"
)

type Page struct {
	URL    *url.URL         `json:"-"`
	PATH   string           `json:"url"`
	Assets assets.Extractor `json:"assets"`
}

func NewPage(URL string) (*Page, error) {
	validURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	return &Page{
		URL:  validURL,
		PATH: fmt.Sprintf("%s%s", validURL.Host, validURL.Path),
	}, nil
}

func (p *Page) setAssets(Assets assets.Extractor) *Page {
	p.Assets = Assets
	return p
}

func (p *Page) AddHTMLContent(content []byte) (*Page, error) {
	newAssets, err := assets.NewAssets(content)
	if err != nil {
		return nil, err
	}
	p.setAssets(newAssets)
	return p, nil
}

func (p Page) String() string {
	jsonString, _ := json.MarshalIndent(p, "", "\t")

	return string(jsonString)
}
