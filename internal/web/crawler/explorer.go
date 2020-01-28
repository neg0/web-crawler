package crawler

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"cuvva/internal/web/page"
	"cuvva/pkg/adapter/html"
	"cuvva/pkg/adapter/http"
)

type ExplorerClient struct {
	HTTPClient http.RequestHandler
	WaitGroup  *sync.WaitGroup
}

func NewExplorerClient() ExplorerClient {
	return ExplorerClient{
		HTTPClient: http.NewClient(),
		WaitGroup:  &sync.WaitGroup{},
	}
}

func (c ExplorerClient) Handle(targetURL string, defaultDepth int) ([]*page.Page, error) {
	if defaultDepth == 0 {
		return nil, errors.New("depth should be greater than zero")
	}

	listOfPages, targetHostName, err := initialHostHomePage(targetURL)
	if err != nil {
		return nil, err
	}

	allTargets := page.NewCollection()
	defaultLength := 0
	for i := 0; i < defaultDepth; i++ {
		prevDefaultLength := defaultLength
		defaultLength = len(allTargets.Pages)
		if i != 0 && prevDefaultLength == defaultLength {
			i = defaultDepth
			fmt.Println("Finished the process")
			continue
		}

		c.WaitGroup.Add(1)
		go func() {
			for _, path := range listOfPages {
				targetPath := fmt.Sprintf("http://%s%s", targetHostName, path.Path)
				fmt.Println("Scanning: ", targetPath)

				resp, err := c.HTTPClient.Get(targetPath)
				if err != nil {
					continue
				}

				htmlReader, _ := html.NewReader(resp.Body())
				links := htmlReader.Find("a")
				listOfPages = findNewLinks(targetHostName, links, listOfPages)

				if !isVisited(allTargets.Pages, path) {
					newPage, _ := page.NewPage(path.String())
					_, _ = newPage.AddHTMLContent(resp.Body())
					allTargets.AddPage(*newPage)
				}
			}
			c.WaitGroup.Done()
		}()
		c.WaitGroup.Wait()
	}
	return allTargets.Pages, nil
}

func initialHostHomePage(targetURL string) ([]*url.URL, string, error) {
	validURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, "", err
	}

	targetHostName := validURL.Host
	targetHostName = strings.Replace(targetHostName, "www.", "", -1)

	homePage, err := url.Parse(validURL.Path)
	if err != nil {
		return nil, targetHostName, err
	}

	return []*url.URL{
		homePage,
	}, targetHostName, nil
}

func findNewLinks(targetHostName string, links []html.ElementDescriber, listOfPages []*url.URL) []*url.URL {
	for _, link := range links {
		trimmedLink := strings.Replace(link.Attr("href"), " ", "", -1)
		validURL, err := url.Parse(trimmedLink)
		if err != nil {
			continue
		}

		if validURL.Host == "" {
			validURL.Host = targetHostName
		}

		if strings.Contains(targetHostName, validURL.Host) &&
			!strings.Contains("@", validURL.Path) &&
			!strings.Contains("#", validURL.Path) {
			if !contains(listOfPages, validURL) {
				listOfPages = append(listOfPages, validURL)
			}
		}
	}
	return listOfPages
}

func contains(a []*url.URL, x *url.URL) bool {
	for _, n := range a {
		if x.RequestURI() == n.RequestURI() {
			return true
		}
	}
	return false
}

func isVisited(visitedTargets []*page.Page, URLPath *url.URL) bool {
	for _, vt := range visitedTargets {
		if vt.URL.Path == URLPath.Path {
			return true
		}
	}
	return false
}
