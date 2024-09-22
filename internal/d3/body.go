package d3

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Body struct {
	size   int
	body   []byte
	reader io.Reader

	Frames Frames
}

func NewBody(f *os.File, size int) (*Body, error) {
	body := make([]byte, size)
	if _, err := io.ReadFull(f, body); err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	b := &Body{
		size:   size,
		body:   body,
		reader: bytes.NewReader(body),
	}

	if err := b.construct(); err != nil {
		return nil, fmt.Errorf("constructing body: %w", err)
	}

	return b, nil
}

func (b *Body) construct() error {
	read := 0
	for read < b.size {
		f, err := NewFrame(b.reader)
		if err != nil {
			return fmt.Errorf("creating frames: %w", err)
		}

		b.Frames = append(b.Frames, f)
		read += f.Size()
	}

	return nil
}

func (b *Body) String() string {
	return fmt.Sprintf("Frame count: %d\nFrames:\n\n%s", len(b.Frames), b.Frames)
}
