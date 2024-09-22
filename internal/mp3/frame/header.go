package frame

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type Header struct {
	bitstream uint32

	id ID
}

const FrameSync uint32 = 0xfff00000

func NewHeader(f *os.File, startAt int) (Header, error) {
	bitstream := make([]byte, 4)
	if _, err := f.ReadAt(bitstream, int64(startAt)); err != nil {
		return Header{}, fmt.Errorf("reading header bytes: %w", err)
	}

	h := Header{bitstream: binary.BigEndian.Uint32(bitstream)}

	if err := h.validate(); err != nil {
		return Header{}, err
	}

	return h, nil
}

func (h *Header) validate() error {
	if !h.isValidSync() {
		return errors.New(fmt.Sprintf("invalid frame sync bits: %032b", h.bitstream))
	}

	h.generateID()

	if h.id == MPEGReservedID || h.id == MPEG25ID {
		return errors.New(fmt.Sprintf("invalid mpeg file id: %s", MPEGReservedID))
	}

	return nil
}

func (h *Header) isValidSync() bool {
	return (h.bitstream & FrameSync) == FrameSync
}

func (h *Header) generateID() {
	h.id = fromFrameHeader(h.bitstream)
}

func (h Header) ID() ID {
	return h.id
}
func (h Header) String() string {
	return fmt.Sprintf("Bistream: %032b\nID: %s\n", h.bitstream, h.id)
}
