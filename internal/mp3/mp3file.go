package mp3

import (
	"fmt"
	"os"

	"github.com/vctaragao/mp3/internal/d3"
	"github.com/vctaragao/mp3/internal/mp3/frame"
)

type File struct {
	file *os.File

	ID3v2  *d3.Tag
	Header frame.Header
	Frame  frame.Frames
}

func New(file *os.File) (File, error) {
	f := File{file: file}

	if err := f.parseID3v2Tag(); err != nil {
		return f, fmt.Errorf("parsing ID3v2 tag: %w", err)
	}

	if err := f.parseFramesHeader(); err != nil {
		return f, fmt.Errorf("parsing frames header: %w", err)
	}

	if err := f.parseFrames(); err != nil {
		return f, fmt.Errorf("parsing frames: %w", err)
	}

	return f, nil
}

func (f *File) parseID3v2Tag() error {
	id3v2Tag, err := d3.New(f.file)
	if err != nil {
		return fmt.Errorf("creating ID3v2 tag: %w", err)
	}

	f.ID3v2 = id3v2Tag

	return nil
}

func (f *File) parseFramesHeader() error {
	id3v2Size := f.ID3v2.ID3Header.Size + d3.HeaderSize

	header, err := frame.NewHeader(f.file, id3v2Size)
	if err != nil {
		return fmt.Errorf("creating header: %w", err)
	}

	f.Header = header

	return nil
}

func (f *File) parseFrames() error {
	return nil
}

func (f *File) ShowID3v2Tag() {
	fmt.Println(f.ID3v2)
}

func (f *File) ShowFramesHeader() {
	fmt.Printf("Frames Header:\n%s", f.Header)
}

func (f *File) ShowID3v2Header() {
	fmt.Printf("ID3v2 Header:\n%s\n", f.ID3v2.ID3Header)
}
