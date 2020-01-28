package assets

type Extractor interface {
	Title() string
	JS() []string
	IMG() []string
	CSS() []string
}

type Assets struct {
	PageTitle           string   `json:"title"`
	JavaScripts         []string `json:"js"`
	Images              []string `json:"images"`
	CascadingStyleSheet []string `json:"css"`
}

func NewAssets(HTMLContent []byte) (*Assets, error) {
	assetsExtractor, err := newAssetsExtractor(HTMLContent)
	if err != nil {
		return nil, err
	}

	return &Assets{
		PageTitle:           assetsExtractor.Title(),
		JavaScripts:         assetsExtractor.JS(),
		Images:              assetsExtractor.IMG(),
		CascadingStyleSheet: assetsExtractor.CSS(),
	}, nil
}

func (a Assets) Title() string {
	return a.PageTitle
}

func (a Assets) JS() []string {
	return a.JavaScripts
}

func (a Assets) IMG() []string {
	return a.Images
}

func (a Assets) CSS() []string {
	return a.CascadingStyleSheet
}
