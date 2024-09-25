package bits

const FrameSync uint32 = 0xfff00000

const ProtectionBit uint32 = 0x00010000 // 0000 0000 0001 0000 0000 0000 0000 0000

func HasProtectionBit(bitstream uint32) bool {
	return bitstream&ProtectionBit == ProtectionBit
}

type SampleFrequency uint32

const (
	sf44100    SampleFrequency = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	sf48000    SampleFrequency = 0x00000400 // 0000 0000 0000 0000 0000 0100 0000 0000
	sf32000    SampleFrequency = 0x00000800 // 0000 0000 0000 0000 0000 1000 0000 0000
	sfReserved SampleFrequency = 0x00000C00 // 0000 0000 0000 0000 0000 1100 0000 0000
)

const PaddingBit uint32 = 0x00000020

func HasPaddingBit(bitstream uint32) bool {
	return bitstream&PaddingBit == PaddingBit
}

const PrivateBit uint32 = 0x00000010

func HasPrivateBit(bitstream uint32) bool {
	return bitstream&PrivateBit == PrivateBit
}

type ChannelMode uint32

const (
	chanStereo ChannelMode = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	chanJoint  ChannelMode = 0x00000040 // 0000 0000 0000 0000 0000 0000 0100 0000
	chanDual   ChannelMode = 0x00000080 // 0000 0000 0000 0000 0000 0000 1000 0000
	chanSingle ChannelMode = 0x000000C0 // 0000 0000 0000 0000 0000 0000 1100 0000
)

type ModeExtension uint32

const (
	band431  ModeExtension = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	band831  ModeExtension = 0x00000010 // 0000 0000 0000 0000 0000 0000 0001 0000
	band1231 ModeExtension = 0x00000020 // 0000 0000 0000 0000 0000 0000 0010 0000
	band1631 ModeExtension = 0x00000030 // 0000 0000 0000 0000 0000 0000 0011 0000
)

const CopyrightBit uint32 = 0x00000008 // 0000 0000 0000 0000 0000 0000 0000 1000

func HasCopyrightBit(bitstream uint32) bool {
	return bitstream&CopyrightBit == CopyrightBit
}

const OriginalBit uint32 = 0x00000006 // 0000 0000 0000 0000 0000 0000 0000 0100

func HasOriginalBit(bitstream uint32) bool {
	return bitstream&OriginalBit == OriginalBit
}

type Emphasis uint32

const (
	none     Emphasis = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	ms       Emphasis = 0x00000001 // 0000 0000 0000 0000 0000 0000 0000 0001
	reserved Emphasis = 0x00000002 // 0000 0000 0000 0000 0000 0000 0000 0010
	ccit     Emphasis = 0x00000003 // 0000 0000 0000 0000 0000 0000 0000 0011
)
