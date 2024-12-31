package frame

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/vctaragao/mp3/internal/mp3/bits"
)

type Header struct {
	bits.BitStream

	id              bits.ID
	layer           bits.Layer
	protectionBit   bool
	bitRate         int
	frequency       bits.Frequency
	paddingBit      bool
	privateBit      bool
	chMode          bits.ChannelMode
	modeExtention   *bits.ModeExtension
	intensityStereo bool
	msStereo        bool
	copyright       bool
	original        bool
	emphasis        bits.Emphasis

	StartAtByte int
	FinishByte  int
}

func NewHeader(f *os.File, startAt int) (Header, error) {
	bitstream := make([]byte, 4)
	if _, err := f.ReadAt(bitstream, int64(startAt)); err != nil {
		return Header{}, fmt.Errorf("reading header bytes: %w", err)
	}

	h := Header{BitStream: bits.BitStream(binary.BigEndian.Uint32(bitstream))}

	if err := h.validate(); err != nil {
		return Header{}, err
	}

	h.StartAtByte = startAt
	h.FinishByte = h.StartAtByte + len(bitstream)

	return h, nil
}

func (h *Header) validate() error {
	if !h.IsValidSync() {
		return errors.New(fmt.Sprintf("invalid frame sync bits: %032b", h.BitStream))
	}

	h.id = h.IDFromFrameHeader()

	if h.id == bits.MPEGReservedID || h.id == bits.MPEG25ID {
		return errors.New(fmt.Sprintf("invalid mpeg file id: %s", bits.MPEGReservedID))
	}

	h.layer = h.ParseLayer()
	h.protectionBit = h.HasProtectionBit()
	h.bitRate = h.ParseBitRate(h.layer, h.id)
	h.frequency = h.ParseSampeFrequency(h.id)
	h.paddingBit = h.HasPaddingBit()
	h.privateBit = h.HasPrivateBit()
	h.chMode = h.ParseChannelMode()

	if h.chMode.IsJointStereo() {
		h.parseModeExtenstion()
	}

	h.copyright = h.HasCopyrightBit()
	h.original = h.HasOriginalBit()
	h.emphasis = h.ParseEmphasis()

	return nil
}

func (h *Header) parseModeExtenstion() {
	if h.layer == bits.Layer3 {
		h.intensityStereo, h.msStereo = h.ParseStereoInfo()
		return
	}

	mdExtension := h.ParseModeExtension()
	h.modeExtention = &mdExtension
}

func (h Header) ID() bits.ID {
	return h.id
}

func (h Header) String() string {
	return fmt.Sprintf("Bistream: %s\nID: %s\nLayer: %s\nProtection bit: %v\nBitRate: %v\nFrequency: %v\nPadding Bit: %v\nPrivate Bit: %v\nChannel Mode: %v\nMode Extension: %v\nItensity Stereo: %v\nMs Stereto: %v\nCopyright: %v\nOriginal: %v\nEmpashis: %v\n",
		h.BitStream,
		h.id,
		h.layer,
		h.protectionBit,
		h.bitRate,
		h.frequency,
		h.paddingBit,
		h.privateBit,
		h.chMode,
		h.modeExtention,
		h.intensityStereo,
		h.msStereo,
		h.copyright,
		h.original,
		h.emphasis,
	)
}

func (h Header) FrameLength() int {
	padding := 0
	if h.paddingBit {
		padding = 1
	}

	if h.layer == bits.Layer1 {
		return (12*h.bitRate/h.frequency.Int() + padding) * 4
	}

	fmt.Println("calculating for layer 3")
	fmt.Println("bitRate", h.bitRate)
	fmt.Println("frequency", h.frequency.Int())
	fmt.Println("padding", padding)

	return 144*h.bitRate*1000/h.frequency.Int() + padding
}
