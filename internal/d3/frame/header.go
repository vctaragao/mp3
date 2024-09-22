package frame

import (
	"bytes"
	"fmt"
	"io"
	"math"
)

const HeaderSize = 10

type Header struct {
	reader io.Reader
	header []byte

	Identifier Identifier
	Size       int
	Flags      [2]byte
}

func NewHeader(reader io.Reader) (Header, error) {
	header := make([]byte, HeaderSize)
	if _, err := io.ReadFull(reader, header); err != nil {
		return Header{}, fmt.Errorf("reading header bytes: %w", err)
	}

	h := Header{
		header: header,
		reader: bytes.NewReader(header),
	}

	if err := h.construct(); err != nil {
		return h, fmt.Errorf("constructing header: %w", err)
	}

	return h, nil
}

func (h *Header) construct() error {
	idBytes := make([]byte, 4)
	if _, err := io.ReadFull(h.reader, idBytes); err != nil {
		return fmt.Errorf("reading identifier: %w", err)
	}

	identifier, err := IdentifierFromString(string(idBytes))
	if err != nil {
		return fmt.Errorf("converting Identifier from string: %w", err)
	}

	h.Identifier = identifier

	size := make([]byte, 4)
	if _, err := io.ReadFull(h.reader, size); err != nil {
		return fmt.Errorf("reading size: %w", err)
	}

	for i, s := range size {
		bitPosition := math.Pow(2, 8*(4-(float64(i)+1)))
		h.Size += int(s) * int(bitPosition)
	}

	flags := [2]byte{}
	if _, err := io.ReadFull(h.reader, flags[:]); err != nil {
		return fmt.Errorf("reading flags: %w", err)
	}

	h.Flags = flags

	return nil
}

func (h Header) String() string {
	return fmt.Sprintf("ID: %s\nSize: %d\nFlags: %04b", h.Identifier, h.Size, h.Flags)
}
