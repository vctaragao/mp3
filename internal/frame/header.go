package frame

import (
	"errors"
	"fmt"
)

type Header struct {
	bitstream uint32
	id        ID
}

const FrameSync uint32 = 0xffe00000

func NewHeader(bitstream uint32) (Header, error) {
	f := Header{
		bitstream: bitstream,
	}

	if err := f.validate(); err != nil {
		return Header{}, err
	}

	return f, nil
}

func (f *Header) validate() error {
	if !f.isValidSync() {
		return errors.New("invalid frame sync bits")
	}

	if f.generateID() == MPEGReservedID {
		return errors.New(fmt.Sprintf("invalid mpeg file id: %s", MPEGReservedID))
	}

	return nil
}

func (f *Header) isValidSync() bool {
	return (f.bitstream & FrameSync) == FrameSync
}

func (f *Header) generateID() ID {
	f.id = fromFrameHeader(f.bitstream)
	return f.id
}

func (f Header) ID() ID {
	return f.id
}
