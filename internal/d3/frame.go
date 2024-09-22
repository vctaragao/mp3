package d3

import (
	"fmt"
	"io"

	"github.com/vctaragao/mp3/internal/d3/frame"
)

type (
	Frame struct {
		header frame.Header
		body   frame.Body
	}

	Frames []Frame
)

func (frames Frames) String() string {
	strFrames := ""

	for _, f := range frames {
		strFrames += fmt.Sprintf("%s\n\n", f)
	}

	return strFrames
}

func NewFrame(reader io.Reader) (Frame, error) {
	f := Frame{}
	header, err := frame.NewHeader(reader)
	if err != nil {
		return f, fmt.Errorf("creating header: %w", err)
	}

	f.header = header

	body, err := frame.NewBody(reader, f.header.Size, f.header.Identifier)
	if err != nil {
		return f, fmt.Errorf("creating body: %w", err)
	}

	f.body = body

	return f, nil
}

func (f Frame) Size() int {
	return f.header.Size + frame.HeaderSize
}

func (f Frame) String() string {
	return fmt.Sprintf("%s\n%s", f.header, f.body)

}
