package html

type Finder interface {
	Find(ElementName string) []ElementDescriber
}

type ElementDescriber interface {
	Attr(string) string
	Text() string
}

type Reader struct {
	Finder
}

func (r Reader) Find(ElementName string) []ElementDescriber {
	return r.Finder.Find(ElementName)
}

func NewReader(HTMLContent []byte) (*Reader, error) {
	goQueryDoc, err := newGoQueryDocument(HTMLContent)
	if err != nil {
		return nil, err
	}

	return &Reader{goQueryDoc}, nil
}
