package decoder

import (
	"fmt"
	"os"
)

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(f *os.File) (mp3File, error) {
	file, err := NewMp3File(f)
	if err != nil {
		return file, fmt.Errorf("creating mp3file: %w", err)
	}

	return file, nil
}
