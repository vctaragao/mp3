package decoder

import (
	"fmt"
	"os"

	"github.com/vctaragao/mp3/internal/d3"
	"github.com/vctaragao/mp3/internal/frame"
)

type mp3File struct {
	file *os.File

	ID3Header *d3.Header
	Frame     frame.Frame
}

func NewMp3File(file *os.File) (mp3File, error) {
	f := mp3File{file: file}

	id3v2Header, err := d3.NewHeader(f.file)
	if err != nil {
		return f, fmt.Errorf("creating ID3v2 header: %w", err)
	}

	f.ID3Header = id3v2Header

	return f, nil
}

func (f *mp3File) ShowHeader() {
	fmt.Println(f.ID3Header)
}
