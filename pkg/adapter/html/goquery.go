package html

import (
	goQueryDirect "github.com/PuerkitoBio/goquery"

	"cuvva/pkg/adapter/html/goquery"
)

type goQueryElement struct {
	goquery.Selection
}

func (gqe goQueryElement) Attr(attr string) string {
	val, exist := gqe.Selection.Attr(attr)
	if !exist {
		return ""
	}
	return val
}

func (gqe goQueryElement) Text() string {
	return gqe.Selection.Text()
}

type goQueryDocument struct {
	goquery.DocumentReader
}

func newGoQueryDocument(HTMLContent []byte) (*goQueryDocument, error) {
	goQueryDoc, err := goquery.NewReader(HTMLContent)
	if err != nil {
		return nil, err
	}

	return &goQueryDocument{DocumentReader: goQueryDoc}, nil
}

func (gqd goQueryDocument) Find(ElementName string) []ElementDescriber {
	var elements []ElementDescriber
	gqd.DocumentReader.Find(ElementName).Each(func(i int, selection *goQueryDirect.Selection) {
		elements = append(elements, goQueryElement{
			Selection: goquery.NewSelection(selection),
		})
	})
	return elements
}
