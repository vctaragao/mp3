package frameheader

import (
	"errors"
	"fmt"
)

type FrameHeader struct {
	bitstream uint32
	id        ID
}

const FrameSync uint32 = 0xffe00000

func NewFrameHeader(bitstream uint32) (FrameHeader, error) {
	f := FrameHeader{
		bitstream: bitstream,
	}

	if err := f.validate(); err != nil {
		return FrameHeader{}, err
	}

	return f, nil
}

func (f *FrameHeader) validate() error {
	if !f.isValidSync() {
		return errors.New("invalid frame sync bits")
	}

	if f.generateID() == MPEGReservedID {
		return errors.New(fmt.Sprintf("invalid mpeg file id: %s", MPEGReservedID))
	}

	return nil
}

func (f *FrameHeader) isValidSync() bool {
	return (f.bitstream & FrameSync) == FrameSync
}

func (f *FrameHeader) generateID() ID {
	f.id = fromFrameHeader(f.bitstream)
	return f.id
}

func (f FrameHeader) ID() ID {
	return f.id
}
