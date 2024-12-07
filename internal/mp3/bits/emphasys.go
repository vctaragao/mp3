package bits

type Emphasis int

const (
	none Emphasis = iota
	ms
	reserved
	ccit

	noneID     BitStream = 0x00000000 // 0000 0000 0000 0000 0000 0000 0000 0000
	msID       BitStream = 0x00000001 // 0000 0000 0000 0000 0000 0000 0000 0001
	reservedID BitStream = 0x00000002 // 0000 0000 0000 0000 0000 0000 0000 0010
	ccitID     BitStream = 0x00000003 // 0000 0000 0000 0000 0000 0000 0000 0011
)

func (e Emphasis) String() string {
	return []string{"None", "MS", "Reserved", "CCIT"}[e]
}

func (b BitStream) ParseEmphasis() Emphasis {
	if b&ccitID == ccitID {
		return ccit
	}

	if b&reservedID == reservedID {
		return reserved
	}

	if b&msID == msID {
		return ms
	}

	return none
}
