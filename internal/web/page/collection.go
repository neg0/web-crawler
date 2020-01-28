package page

import "cuvva/pkg/sitemap"

type Collection struct {
	Pages []*Page
}

func NewCollection() Collection {
	return Collection{Pages: []*Page{}}
}

func NewPageCollectionFromSiteMap(siteMap sitemap.Reader) (*Collection, error) {
	var pageCollection []*Page
	for _, path := range siteMap.URLs() {
		pageBuilder, err := NewPage(path)
		if err != nil {
			return nil, err
		}
		pageCollection = append(pageCollection, pageBuilder)
	}
	return &Collection{Pages: pageCollection}, nil
}

func (c *Collection) AddPage(page Page) Collection {
	c.Pages = append(c.Pages, &page)
	return *c
}
