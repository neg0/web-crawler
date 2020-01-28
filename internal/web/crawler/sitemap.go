package crawler

import (
	"fmt"
	"net/url"
	"sync"

	"cuvva/internal/web/page"
	"cuvva/pkg/adapter/http"
	"cuvva/pkg/sitemap"
)

type SiteMapCrawler struct {
	HTTPClient http.RequestHandler
	WaitGroup  *sync.WaitGroup
}

func NewSiteMapCrawler() SiteMapCrawler {
	return SiteMapCrawler{
		HTTPClient: http.NewClient(),
		WaitGroup:  &sync.WaitGroup{},
	}
}

func (smc SiteMapCrawler) Handle(TargetURL string) ([]*page.Page, error) {
	validURL, validationErr := url.Parse(TargetURL)
	if validationErr != nil {
		return nil, validationErr
	}

	resp, err := smc.HTTPClient.Get(fmt.Sprintf("https://%s/sitemap.xml", validURL.Host))
	if err != nil {
		return nil, err
	}

	siteMap, smErr := sitemap.NewSiteMap(resp.Body())
	if smErr != nil {
		return nil, smErr
	}

	pageCollection, pcErr := page.NewPageCollectionFromSiteMap(siteMap)
	if pcErr != nil {
		return nil, pcErr
	}

	smc.WaitGroup.Add(1)
	go func() {
		for _, c := range pageCollection.Pages {
			fmt.Println("Scanning:", c.URL.Path)
			resp, err := smc.HTTPClient.Get(c.URL.String())
			if err != nil {
				continue
			}
			_, _ = c.AddHTMLContent(resp.Body())
		}
		smc.WaitGroup.Done()
	}()
	smc.WaitGroup.Wait()

	return pageCollection.Pages, nil
}
