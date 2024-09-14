package decoder

import (
	"fmt"
	"io"
	"os"
)

type Decoder struct{}

func NewDecoder() Decoder {
	return Decoder{}
}

func (d *Decoder) Decode(f *os.File) {
	headerBytes := make([]byte, 256)

	_, err := io.ReadFull(f, headerBytes)
	if err != nil {
		panic(err)
	}

	// TODO: deal with ID3v2 header First 10 bytes (byte 6-10 indicate the full ID3v2 header size)
	// https://id3.org/id3v2.3.0
	// ID3v2/file identifier   "ID3"
	// ID3v2 version           $03 00
	// ID3v2 flags             %abc00000
	// ID3v2 size              4 * %0xxxxxxx

	for i, b := range headerBytes {
		fmt.Printf("%d - %b : %d : %s\n", i+1, b, b, string(b))
	}
}
