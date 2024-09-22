package d3

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

const (
	HeaderSize = 10
	MinVersion = 3
	Identifier = "ID3"
)

type Version int

func (v Version) String() string {
	return fmt.Sprintf("ID3v2.%d.x", v)
}

type Header struct {
	header []byte
	reader io.Reader

	Size       int
	Flags      byte
	Identifier string
	Version    Version
}

func NewHeader(f *os.File) (*Header, error) {
	header := make([]byte, HeaderSize)

	if _, err := io.ReadFull(f, header); err != nil {
		return &Header{}, fmt.Errorf("reading first 10 bytes of file: %w", err)
	}

	h := &Header{
		header: header,
		reader: bytes.NewReader(header),
	}

	if err := h.construct(); err != nil {
		return h, fmt.Errorf("constructing ID3v2 header: %w", err)
	}

	if err := h.validate(); err != nil {
		return h, fmt.Errorf("validating ID3v2 header: %w", err)
	}

	return h, nil
}

func (h *Header) construct() error {
	identifier := make([]byte, 3)
	if _, err := io.ReadFull(h.reader, identifier); err != nil {
		return err
	}

	h.Identifier = string(identifier)

	version := make([]byte, 2)
	if _, err := io.ReadFull(h.reader, version); err != nil {
		return err
	}

	h.Version = Version(version[0])

	flags := make([]byte, 1)
	if _, err := io.ReadFull(h.reader, flags); err != nil {
		return err
	}

	h.Flags = flags[0]

	size := make([]byte, 4)
	if _, err := io.ReadFull(h.reader, size); err != nil {
		return err
	}

	for i, s := range size {
		bitPlacement := math.Pow(2, float64(7*(4-(i+1))))
		h.Size += int(s) * int(bitPlacement)
	}

	return nil
}

func (h *Header) validate() error {
	if h.Identifier != Identifier {
		return errors.New("unsuported, not an ID3v2 header")
	}

	if h.Version != MinVersion {
		return errors.New("unsoported ID3v2 Header version")
	}

	if h.Size == 0 {
		return errors.New("empty ID3v2 header")
	}

	return nil
}

func (h *Header) String() string {
	return fmt.Sprintf("Identifier: %s\nVersion: %s\nFlags: %08b\nSize: %d", h.Identifier, h.Version, h.Flags, h.Size)
}
