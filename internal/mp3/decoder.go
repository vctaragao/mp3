package mp3

import (
	"fmt"
	"os"
)

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(f *os.File) (File, error) {
	file, err := New(f)
	if err != nil {
		return file, fmt.Errorf("creating mp3file: %w", err)
	}

	return file, nil
}
