package mp3

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	"github.com/vctaragao/mp3/internal/d3"
	"github.com/vctaragao/mp3/internal/mp3/bits"
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

// TODO: Parse Frames
func (f *File) parseFrames() error {
	frameData := make([]byte, f.Header.FrameLength())
	if _, err := f.file.ReadAt(frameData, int64(f.Header.FinishByte)); err != nil {
		return fmt.Errorf("reading first frame data after header: %w", err)
	}

	byiteIndex := 0
	dataReader := bytes.NewReader(frameData)

	for {
		bitstream := make([]byte, 4)
		n, err := dataReader.Read(bitstream)
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return fmt.Errorf("reading frame data: %w", err)
		}
		byiteIndex += n

		b := binary.BigEndian.Uint32(bitstream)
		if bits.BitStream(b)&bits.FrameSync == bits.FrameSync {
			// TODO: It shows at 1044 position, just before the padding position
			// The length of this frame is 1045
			//
			// TODO: It probably meas that as I read the frame body data I will need
			// to be checking if I arrived at another frame header
			fmt.Printf("Index: %d, %b\n", byiteIndex, b)
		}
	}
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
