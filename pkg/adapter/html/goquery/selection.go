package goquery

import "github.com/PuerkitoBio/goquery"

type Selector interface {
	Text() string
	Attr(attrName string) (val string, exists bool)
	Each(f func(int, *goquery.Selection)) *goquery.Selection
}

type Selection struct {
	Selector
}

func NewSelection(Selector Selector) Selection {
	return Selection{Selector: Selector}
}

func (s Selection) Text() string {
	return s.Selector.Text()
}

func (s Selection) Attr(attrName string) (val string, exists bool) {
	return s.Selector.Attr(attrName)
}

func (s Selection) Each(f func(int, *goquery.Selection)) *goquery.Selection {
	return s.Selector.Each(f)
}
