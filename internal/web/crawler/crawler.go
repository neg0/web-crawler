package crawler

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"cuvva/internal/web/page"
	"cuvva/pkg/adapter/http"
)

type TimeOutErr struct{}

func (toe TimeOutErr) Error() string {
	return "timed out"
}

func Handle(TargetURL string) ([]*page.Page, error) {
	parsedURL, err := url.ParseRequestURI(TargetURL)
	if err != nil {
		return nil, err
	}

	isSiteMapAvailable := false
	httpClient := http.NewClient()
	targetSiteMap := fmt.Sprintf("%s://%s/sitemap.xml", parsedURL.Scheme, parsedURL.Host)
	_, errRq := httpClient.Get(targetSiteMap)
	if errRq == nil {
		isSiteMapAvailable = true
	}

	respCh := make(chan []*page.Page)
	timeOutDefault, _ := strconv.ParseInt(os.Getenv("DEFAULT_TIMEOUT_IN_SECONDS"), 0, 64)
	if timeOutDefault < 1 {
		timeOutDefault = 3
	}

	if !isSiteMapAvailable {
		explorer := NewExplorerClient()
		go func() {
			resp, _ := explorer.Handle(fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host), 4)
			respCh <- resp
			close(respCh)
		}()

		for {
			select {
			case res := <-respCh:
				return res, nil
			case <-time.After(time.Second * time.Duration(timeOutDefault)):
				return nil, TimeOutErr{}
			}
		}
	}

	siteMapCrawler := NewSiteMapCrawler()
	go func() {
		resp, _ := siteMapCrawler.Handle(fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host))
		respCh <- resp
		close(respCh)
	}()

	for {
		select {
		case res := <-respCh:
			return res, nil
		case <-time.After(time.Second * time.Duration(timeOutDefault)):
			return nil, TimeOutErr{}
		}
	}
}
