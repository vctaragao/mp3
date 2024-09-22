package decoder

import (
	"fmt"
	"os"

	"github.com/vctaragao/mp3/internal/d3"
	"github.com/vctaragao/mp3/internal/frame"
)

type mp3File struct {
	file *os.File

	ID3v2 *d3.Tag
	Frame frame.Frame
}

func NewMp3File(file *os.File) (mp3File, error) {
	f := mp3File{file: file}

	id3v2Tag, err := d3.New(f.file)
	if err != nil {
		return f, fmt.Errorf("creating ID3v2 tag: %w", err)
	}

	f.ID3v2 = id3v2Tag

	return f, nil
}
func (f *mp3File) ShowID3v2Tag() {
	fmt.Println(f.ID3v2)
}

func (f *mp3File) ShowHeader() {
	fmt.Println(f.ID3v2.ID3Header)
}
