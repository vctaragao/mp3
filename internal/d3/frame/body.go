package frame

import (
	"bytes"
	"fmt"
	"io"
)

type Body struct {
	body   []byte
	reader io.Reader

	strContent string
}

func NewBody(reader io.Reader, size int, id Identifier) (Body, error) {
	body := make([]byte, size)
	if _, err := io.ReadFull(reader, body); err != nil {
		return Body{}, fmt.Errorf("reading body: %w", err)
	}

	b := Body{
		body:   body,
		reader: bytes.NewReader(body),
	}

	switch id {
	case TIT2, TYER, TENC, TBPM, TCON, TPE1:
		b.constructText()
	}

	return b, nil
}

func (b *Body) constructText() {
	b.strContent = string(b.body)
}

func (b Body) String() string {
	return fmt.Sprintf("Content: %s", b.strContent)
}
