package bits

type ModeExtension int

const (
	b431 ModeExtension = iota
	b831
	b1231
	b1631

	band431  BitStream = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	band831  BitStream = 0x00000010 // 0000 0000 0000 0000 0000 0000 0001 0000
	band1231 BitStream = 0x00000020 // 0000 0000 0000 0000 0000 0000 0010 0000
	band1631 BitStream = 0x00000030 // 0000 0000 0000 0000 0000 0000 0011 0000
)

func (b BitStream) ParseModeExtension() ModeExtension {
	if b&band1631 == band1631 {
		return b1631
	}

	if b&band1231 == band1231 {
		return b1231
	}

	if b&band831 == band831 {
		return b831
	}

	return b431
}

func (b BitStream) ParseStereoInfo() (bool, bool) {
	if b&band1631 == band1631 {
		return true, true
	}

	if b&band1231 == band1231 {
		return false, true
	}

	if b&band831 == band831 {
		return true, false
	}

	return false, false
}
