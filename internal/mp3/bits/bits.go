package bits

import "fmt"

type BitStream uint32

const (
	FrameSync     BitStream = 0xfff00000 // 1111 1111 1111 0000 0000 0000 0000 0000
	ProtectionBit BitStream = 0x00010000 // 0000 0000 0001 0000 0000 0000 0000 0000
	PaddingBit    BitStream = 0x00000020 // 0000 0000 0000 0000 0000 0010 0000 0000
	PrivateBit    BitStream = 0x00000010 // 0000 0000 0000 0000 0000 0001 0000 0000
	CopyrightBit  BitStream = 0x00000008 // 0000 0000 0000 0000 0000 0000 0000 1000
	OriginalBit   BitStream = 0x00000006 // 0000 0000 0000 0000 0000 0000 0000 0100
)

func (b BitStream) String() string {
	return fmt.Sprintf("%032b", b)
}

func (b BitStream) IsValidSync() bool {
	return b&FrameSync == FrameSync
}

func (b BitStream) HasProtectionBit() bool {
	return b&ProtectionBit == ProtectionBit
}

func (b BitStream) HasPaddingBit() bool {
	return b&PaddingBit == PaddingBit
}

func (b BitStream) HasPrivateBit() bool {
	return b&PrivateBit == PrivateBit
}

func (b BitStream) HasCopyrightBit() bool {
	return b&CopyrightBit == CopyrightBit
}

func (b BitStream) HasOriginalBit() bool {
	return b&OriginalBit == OriginalBit
}
