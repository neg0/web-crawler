package sitemap

import "encoding/xml"

type Reader interface {
	URLs() []string
}

type URL struct {
	Loc string `xml:"loc"`
}

type SiteMap struct {
	XMLName xml.Name `xml:"-"`
	URL     []URL    `xml:"url"`
}

func (sm SiteMap) URLs() []string {
	var collection []string
	for _, URL := range sm.URL {
		collection = append(collection, URL.Loc)
	}
	return collection
}

func NewSiteMap(XMLContent []byte) (*SiteMap, error) {
	var siteMap *SiteMap
	err := xml.Unmarshal(XMLContent, &siteMap)
	if err != nil {
		return nil, err
	}
	return siteMap, nil
}
