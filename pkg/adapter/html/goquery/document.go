package goquery

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

type DocumentReader interface {
	Find(string) *goquery.Selection
}

type Document struct {
	DocumentReader
}

func NewReader(HTMLContent []byte) (*Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(HTMLContent))
	if err != nil {
		return nil, err
	}

	return &Document{DocumentReader: doc}, nil
}

func (r Document) Find(selector string) *goquery.Selection {
	return r.DocumentReader.Find(selector)
}
