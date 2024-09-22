package decoder

import (
	"fmt"
	"os"

	"github.com/vctaragao/mp3/internal/mp3"
)

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(f *os.File) (mp3.File, error) {
	file, err := mp3.New(f)
	if err != nil {
		return file, fmt.Errorf("creating mp3file: %w", err)
	}

	return file, nil
}
