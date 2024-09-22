package d3

import (
	"fmt"
	"os"
)

type Tag struct {
	ID3Header *Header
	ID3Body   *Body
}

func New(f *os.File) (*Tag, error) {
	t := &Tag{}

	header, err := NewHeader(f)
	if err != nil {
		return t, fmt.Errorf("creating ID3v2 tag: %w", err)
	}

	t.ID3Header = header

	body, err := NewBody(f, t.ID3Header.Size)
	if err != nil {
		return t, fmt.Errorf("creating ID3v2 header: %w", err)
	}

	t.ID3Body = body

	return t, nil

}

func (t *Tag) String() string {
	return fmt.Sprintf("Tag Header:\n%s\n\nTag Body: \n%s", t.ID3Header, t.ID3Body)
}
